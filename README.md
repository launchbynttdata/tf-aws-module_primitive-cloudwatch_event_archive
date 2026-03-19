# TF AWS Module Primitive - CloudWatch Event Archive

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![License: CC BY-NC-ND 4.0](https://img.shields.io/badge/License-CC_BY--NC--ND_4.0-lightgrey.svg)](https://creativecommons.org/licenses/by-nc-nd/4.0/)

## Overview

This Terraform module creates an [AWS EventBridge (CloudWatch Events) event archive](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_event_archive) for long-term archival and replay of events from an event bus.

## Pre-Commit Hooks

[.pre-commit-config.yaml](.pre-commit-config.yaml) defines pre-commit hooks for Terraform, Go, and common linting. The `commitlint` hook enforces conventional commit format. The `detect-secrets-hook` prevents new secrets from being introduced. See [pre-commit](https://pre-commit.com/#install) for installation. Install the commit-msg hook manually:

```
pre-commit install --hook-type commit-msg
```

## Usage

See [examples/complete](examples/complete) for a full working example.

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.5 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 6.1 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | 6.37.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_cloudwatch_event_archive.archive](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_event_archive) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_description"></a> [description](#input\_description) | Description for the archive. Maximum 512 characters. | `string` | `null` | no |
| <a name="input_event_pattern"></a> [event\_pattern](#input\_event\_pattern) | Event pattern to filter events sent to the archive. JSON string. Archives all events if not specified. | `string` | `null` | no |
| <a name="input_event_source_arn"></a> [event\_source\_arn](#input\_event\_source\_arn) | ARN of the event bus associated with the archive. Events from this bus are archived. | `string` | n/a | yes |
| <a name="input_kms_key_identifier"></a> [kms\_key\_identifier](#input\_kms\_key\_identifier) | The ARN, Key ID, or alias of the KMS key EventBridge uses to encrypt the archive. Omit to use the AWS owned key. | `string` | `null` | no |
| <a name="input_name"></a> [name](#input\_name) | Name of the EventBridge event archive. Must be unique per account and region. Maximum 48 characters. | `string` | n/a | yes |
| <a name="input_retention_days"></a> [retention\_days](#input\_retention\_days) | Maximum number of days to retain events in the archive. Omit for indefinite retention. | `number` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_arn"></a> [arn](#output\_arn) | The ARN of the archive. |
| <a name="output_description"></a> [description](#output\_description) | The description of the archive. |
| <a name="output_event_pattern"></a> [event\_pattern](#output\_event\_pattern) | The event pattern used to filter events sent to the archive. |
| <a name="output_event_source_arn"></a> [event\_source\_arn](#output\_event\_source\_arn) | The ARN of the event bus associated with the archive. |
| <a name="output_id"></a> [id](#output\_id) | The ID of the archive (same as the name). |
| <a name="output_kms_key_identifier"></a> [kms\_key\_identifier](#output\_kms\_key\_identifier) | The ARN or identifier of the KMS key used to encrypt the archive. |
| <a name="output_name"></a> [name](#output\_name) | The name of the archive. |
| <a name="output_retention_days"></a> [retention\_days](#output\_retention\_days) | The number of days events are retained in the archive. |
<!-- END_TF_DOCS -->
