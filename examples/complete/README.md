# Complete Example

This example creates an EventBridge event bus and an event archive that archives events from that bus.

## Usage

```hcl
data "aws_region" "current" {}

module "resource_names" {
  source  = "terraform.registry.launch.nttdata.com/module_library/resource_name/launch"
  version = "~> 2.0"

  for_each = var.resource_names_map

  logical_product_family  = var.logical_product_family
  logical_product_service = var.logical_product_service
  class_env               = var.class_env
  instance_env            = var.instance_env
  instance_resource      = var.instance_resource
  cloud_resource_type    = each.value.name
  maximum_length         = each.value.max_length

  region = join("", split("-", data.aws_region.current.name))
}

resource "aws_cloudwatch_event_bus" "bus" {
  name = module.resource_names["event_bus"].standard
  tags = {}
}

module "archive" {
  source = "../.."

  name             = module.resource_names["event_archive"].minimal_random_suffix
  event_source_arn = aws_cloudwatch_event_bus.bus.arn
  description      = var.description
  event_pattern    = var.event_pattern
  retention_days   = var.retention_days
}
```

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_resource_names_map"></a> [resource_names_map](#input_resource_names_map) | Map of resource names for the resource naming module. | `map(object({ name = string, max_length = number }))` | n/a | yes |
| <a name="input_logical_product_family"></a> [logical_product_family](#input_logical_product_family) | Logical product family for resource naming. | `string` | n/a | yes |
| <a name="input_logical_product_service"></a> [logical_product_service](#input_logical_product_service) | Logical product service for resource naming. | `string` | n/a | yes |
| <a name="input_class_env"></a> [class_env](#input_class_env) | Class environment for resource naming. | `string` | n/a | yes |
| <a name="input_instance_env"></a> [instance_env](#input_instance_env) | Instance environment number for resource naming. | `number` | n/a | yes |
| <a name="input_instance_resource"></a> [instance_resource](#input_instance_resource) | Instance resource number for resource naming. | `number` | n/a | yes |
| <a name="input_description"></a> [description](#input_description) | Description for the archive. Maximum 512 characters. | `string` | `null` | no |
| <a name="input_event_pattern"></a> [event_pattern](#input_event_pattern) | Event pattern to filter events sent to the archive. JSON string. Archives all events if not specified. | `string` | `null` | no |
| <a name="input_retention_days"></a> [retention_days](#input_retention_days) | Maximum number of days to retain events in the archive. Omit for indefinite retention. | `number` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_region"></a> [region](#output_region) | The AWS region where resources are deployed. |
| <a name="output_id"></a> [id](#output_id) | The ID of the archive. |
| <a name="output_arn"></a> [arn](#output_arn) | The ARN of the archive. |
| <a name="output_name"></a> [name](#output_name) | The name of the archive. |
| <a name="output_event_source_arn"></a> [event_source_arn](#output_event_source_arn) | The ARN of the event bus associated with the archive. |
| <a name="output_description"></a> [description](#output_description) | The description of the archive. |
| <a name="output_event_pattern"></a> [event_pattern](#output_event_pattern) | The event pattern used to filter events sent to the archive. |
| <a name="output_retention_days"></a> [retention_days](#output_retention_days) | The number of days events are retained in the archive. |
| <a name="output_kms_key_identifier"></a> [kms_key_identifier](#output_kms_key_identifier) | The ARN or identifier of the KMS key used to encrypt the archive. |

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

| Name | Source | Version |
|------|--------|---------|
| <a name="module_archive"></a> [archive](#module\_archive) | ../.. | n/a |
| <a name="module_resource_names"></a> [resource\_names](#module\_resource\_names) | terraform.registry.launch.nttdata.com/module_library/resource_name/launch | ~> 2.0 |

## Resources

| Name | Type |
|------|------|
| [aws_cloudwatch_event_bus.bus](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_event_bus) | resource |
| [aws_region.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/region) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_class_env"></a> [class\_env](#input\_class\_env) | Class environment for resource naming. | `string` | n/a | yes |
| <a name="input_description"></a> [description](#input\_description) | Description for the archive. Maximum 512 characters. | `string` | `null` | no |
| <a name="input_event_pattern"></a> [event\_pattern](#input\_event\_pattern) | Event pattern to filter events sent to the archive. JSON string. Archives all events if not specified. | `string` | `null` | no |
| <a name="input_instance_env"></a> [instance\_env](#input\_instance\_env) | Instance environment number for resource naming. | `number` | n/a | yes |
| <a name="input_instance_resource"></a> [instance\_resource](#input\_instance\_resource) | Instance resource number for resource naming. | `number` | n/a | yes |
| <a name="input_logical_product_family"></a> [logical\_product\_family](#input\_logical\_product\_family) | Logical product family for resource naming. | `string` | n/a | yes |
| <a name="input_logical_product_service"></a> [logical\_product\_service](#input\_logical\_product\_service) | Logical product service for resource naming. | `string` | n/a | yes |
| <a name="input_resource_names_map"></a> [resource\_names\_map](#input\_resource\_names\_map) | Map of resource names for the resource naming module. | <pre>map(object({<br/>    name       = string<br/>    max_length = number<br/>  }))</pre> | n/a | yes |
| <a name="input_retention_days"></a> [retention\_days](#input\_retention\_days) | Maximum number of days to retain events in the archive. Omit for indefinite retention. | `number` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_arn"></a> [arn](#output\_arn) | The ARN of the archive. |
| <a name="output_description"></a> [description](#output\_description) | The description of the archive. |
| <a name="output_event_pattern"></a> [event\_pattern](#output\_event\_pattern) | The event pattern used to filter events sent to the archive. |
| <a name="output_event_source_arn"></a> [event\_source\_arn](#output\_event\_source\_arn) | The ARN of the event bus associated with the archive. |
| <a name="output_id"></a> [id](#output\_id) | The ID of the archive. |
| <a name="output_kms_key_identifier"></a> [kms\_key\_identifier](#output\_kms\_key\_identifier) | The ARN or identifier of the KMS key used to encrypt the archive. |
| <a name="output_name"></a> [name](#output\_name) | The name of the archive. |
| <a name="output_region"></a> [region](#output\_region) | The AWS region where resources are deployed. |
| <a name="output_retention_days"></a> [retention\_days](#output\_retention\_days) | The number of days events are retained in the archive. |
<!-- END_TF_DOCS -->
