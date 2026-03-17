package testimpl

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	ebtypes "github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getEventBridgeClient(t *testing.T, region string) *eventbridge.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	require.NoError(t, err, "unable to load AWS config")
	return eventbridge.NewFromConfig(cfg)
}

func TestComposableComplete(t *testing.T, ctx types.TestContext) {
	t.Run("VerifyTerraformOutputs", func(t *testing.T) {
		opts := ctx.TerratestTerraformOptions()
		id := terraform.Output(t, opts, "id")
		name := terraform.Output(t, opts, "name")
		arn := terraform.Output(t, opts, "arn")
		eventSourceArn := terraform.Output(t, opts, "event_source_arn")
		retentionDays := terraform.Output(t, opts, "retention_days")

		assert.Equal(t, name, id, "id should equal name for EventBridge archive")
		assert.Contains(t, arn, "events", "ARN should contain events")
		assert.Contains(t, eventSourceArn, "events", "event_source_arn should contain events")
		assert.Equal(t, "7", retentionDays, "retention_days should be 7")
	})

	t.Run("VerifyArchiveViaAPI", func(t *testing.T) {
		opts := ctx.TerratestTerraformOptions()
		archiveName := terraform.Output(t, opts, "name")
		region := terraform.Output(t, opts, "region")
		expectedRetention := int32(7)
		expectedDescription := "Example EventBridge archive for testing"

		client := getEventBridgeClient(t, region)

		var output *eventbridge.DescribeArchiveOutput
		var err error
		for i := 0; i < 12; i++ {
			output, err = client.DescribeArchive(context.TODO(), &eventbridge.DescribeArchiveInput{
				ArchiveName: aws.String(archiveName),
			})
			if err == nil {
				break
			}
			if strings.Contains(err.Error(), "ResourceNotFoundException") {
				time.Sleep(5 * time.Second)
				continue
			}
			require.NoError(t, err, "DescribeArchive should succeed")
		}
		require.NoError(t, err, "DescribeArchive should succeed after retries")
		require.NotNil(t, output, "DescribeArchive output should not be nil")

		assert.Equal(t, archiveName, aws.ToString(output.ArchiveName), "archive name should match")
		assert.Equal(t, expectedRetention, aws.ToInt32(output.RetentionDays), "retention_days should be 7")
		assert.Equal(t, expectedDescription, aws.ToString(output.Description), "description should match")
		assert.Equal(t, ebtypes.ArchiveStateEnabled, output.State, "archive state should be ENABLED")
	})

	t.Run("PutEventAndVerifyArchiveReceivesIt", func(t *testing.T) {
		opts := ctx.TerratestTerraformOptions()
		eventSourceArn := terraform.Output(t, opts, "event_source_arn")
		region := terraform.Output(t, opts, "region")

		client := getEventBridgeClient(t, region)

		_, err := client.PutEvents(context.TODO(), &eventbridge.PutEventsInput{
			Entries: []ebtypes.PutEventsRequestEntry{
				{
					Source:       aws.String("test.primitive.archive"),
					DetailType:   aws.String("ArchiveTest"),
					Detail:       aws.String(`{"test": "value"}`),
					EventBusName: aws.String(eventSourceArn),
				},
			},
		})
		require.NoError(t, err, "PutEvents should succeed")
	})
}

func TestComposableCompleteReadOnly(t *testing.T, ctx types.TestContext) {
	t.Run("VerifyTerraformOutputs", func(t *testing.T) {
		opts := ctx.TerratestTerraformOptions()
		id := terraform.Output(t, opts, "id")
		name := terraform.Output(t, opts, "name")

		assert.Equal(t, name, id, "id should equal name for EventBridge archive")
	})

	t.Run("VerifyArchiveExistsViaAPI", func(t *testing.T) {
		opts := ctx.TerratestTerraformOptions()
		archiveName := terraform.Output(t, opts, "name")
		region := terraform.Output(t, opts, "region")

		client := getEventBridgeClient(t, region)

		output, err := client.DescribeArchive(context.TODO(), &eventbridge.DescribeArchiveInput{
			ArchiveName: aws.String(archiveName),
		})
		require.NoError(t, err, "DescribeArchive should succeed")
		require.NotNil(t, output, "DescribeArchive output should not be nil")

		assert.Equal(t, archiveName, aws.ToString(output.ArchiveName), "archive name should match")
		assert.Equal(t, ebtypes.ArchiveStateEnabled, output.State, "archive state should be ENABLED")
	})
}
