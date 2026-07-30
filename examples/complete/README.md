# Named Value for APIM Service

Please set a provider block with the following, to avoid soft-deletes of the APIM instance which can cause problems with the tests
```
provider "azurerm" {
  features {
    api_management {
      purge_soft_delete_on_destroy = true
      recover_soft_deleted         = true
    }
  }
}
```


<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.0 |
| <a name="requirement_azurerm"></a> [azurerm](#requirement\_azurerm) | ~>3.117 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_apim"></a> [apim](#module\_apim) | terraform.registry.launch.nttdata.com/module_primitive/api_management/azurerm | ~> 1.0 |
| <a name="module_deployment_role_assignment"></a> [deployment\_role\_assignment](#module\_deployment\_role\_assignment) | terraform.registry.launch.nttdata.com/module_primitive/role_assignment/azurerm | ~> 1.0 |
| <a name="module_key_vault"></a> [key\_vault](#module\_key\_vault) | terraform.registry.launch.nttdata.com/module_primitive/key_vault/azurerm | ~> 2.0 |
| <a name="module_key_vault_secret"></a> [key\_vault\_secret](#module\_key\_vault\_secret) | terraform.registry.launch.nttdata.com/module_primitive/key_vault_secret/azurerm | ~> 1.0 |
| <a name="module_named_value"></a> [named\_value](#module\_named\_value) | ../.. | n/a |
| <a name="module_resource_group"></a> [resource\_group](#module\_resource\_group) | terraform.registry.launch.nttdata.com/module_primitive/resource_group/azurerm | ~> 1.0 |
| <a name="module_resource_names"></a> [resource\_names](#module\_resource\_names) | terraform.registry.launch.nttdata.com/module_library/resource_name/launch | ~> 2.0 |
| <a name="module_secret_named_value"></a> [secret\_named\_value](#module\_secret\_named\_value) | ../.. | n/a |
| <a name="module_uai_role_assignment"></a> [uai\_role\_assignment](#module\_uai\_role\_assignment) | terraform.registry.launch.nttdata.com/module_primitive/role_assignment/azurerm | ~> 1.0 |
| <a name="module_user_assigned_identity"></a> [user\_assigned\_identity](#module\_user\_assigned\_identity) | terraform.registry.launch.nttdata.com/module_primitive/user_managed_identity/azurerm | ~> 1.0 |

## Resources

| Name | Type |
|------|------|
| [azurerm_client_config.client](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/data-sources/client_config) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_environment"></a> [environment](#input\_environment) | Environment in which the resource should be provisioned like dev, qa, prod etc. | `string` | `"dev"` | no |
| <a name="input_environment_number"></a> [environment\_number](#input\_environment\_number) | The environment count for the respective environment. Defaults to 000. Increments in value of 1 | `string` | `"000"` | no |
| <a name="input_key_vault_secret_name"></a> [key\_vault\_secret\_name](#input\_key\_vault\_secret\_name) | The name of the Key Vault Secret to be created. | `string` | `"example-secret"` | no |
| <a name="input_key_vault_secret_value"></a> [key\_vault\_secret\_value](#input\_key\_vault\_secret\_value) | The value of the Key Vault Secret to be created. | `string` | `"example-value"` | no |
| <a name="input_product_family"></a> [product\_family](#input\_product\_family) | (Required) Name of the product family for which the resource is created.<br/>    Example: org\_name, department\_name. | `string` | `"dso"` | no |
| <a name="input_product_service"></a> [product\_service](#input\_product\_service) | (Required) Name of the product service for which the resource is created.<br/>    For example, backend, frontend, middleware etc. | `string` | `"apim"` | no |
| <a name="input_public_network_access_enabled"></a> [public\_network\_access\_enabled](#input\_public\_network\_access\_enabled) | Should the API Management Service be accessible from the public internet?<br/>    This option is applicable only to the Management plane, not the API gateway or Developer portal.<br/>    It is required to be true on the creation.<br/>    For sku=Developer/Premium and network\_type=Internal, it must be true.<br/>    It can only be set to false if there is at least one approve private endpoint connection. | `bool` | `true` | no |
| <a name="input_publisher_email"></a> [publisher\_email](#input\_publisher\_email) | The email of publisher/company. | `string` | `"launchdso@nttdata.com"` | no |
| <a name="input_publisher_name"></a> [publisher\_name](#input\_publisher\_name) | The name of publisher/company. | `string` | `"launchdso"` | no |
| <a name="input_region"></a> [region](#input\_region) | Azure Region in which the infra needs to be provisioned | `string` | `"eastus"` | no |
| <a name="input_resource_names_map"></a> [resource\_names\_map](#input\_resource\_names\_map) | A map of key to resource\_name that will be used by tf-launch-module\_library-resource\_name to generate resource names | <pre>map(object(<br/>    {<br/>      name       = string<br/>      max_length = optional(number, 60)<br/>    }<br/>  ))</pre> | <pre>{<br/>  "api_management": {<br/>    "max_length": 50,<br/>    "name": "apim"<br/>  },<br/>  "key_vault": {<br/>    "max_length": 24,<br/>    "name": "kv"<br/>  },<br/>  "resource_group": {<br/>    "max_length": 50,<br/>    "name": "rg"<br/>  },<br/>  "user_assigned_identity": {<br/>    "max_length": 50,<br/>    "name": "uai"<br/>  }<br/>}</pre> | no |
| <a name="input_resource_number"></a> [resource\_number](#input\_resource\_number) | The resource count for the respective resource. Defaults to 000. Increments in value of 1 | `string` | `"000"` | no |
| <a name="input_sku_name"></a> [sku\_name](#input\_sku\_name) | String consisting of two parts separated by an underscore. The fist part is the name, valid values include: Developer,<br/>    Basic, Standard and Premium. The second part is the capacity. Default is Consumption\_0. | `string` | `"Consumption_0"` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to assign to the resource. | `map(string)` | `{}` | no |
| <a name="input_use_service_principal"></a> [use\_service\_principal](#input\_use\_service\_principal) | Set to false when running locally without a service principal | `bool` | `true` | no |
| <a name="input_virtual_network_type"></a> [virtual\_network\_type](#input\_virtual\_network\_type) | The type of virtual network you want to use, valid values include: None, External, Internal.<br/>    External and Internal are only supported in the SKUs - Premium and Developer | `string` | `"None"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_api_management_id"></a> [api\_management\_id](#output\_api\_management\_id) | The ID of the API Management Service |
| <a name="output_api_management_name"></a> [api\_management\_name](#output\_api\_management\_name) | The name of the API Management Service |
| <a name="output_named_value_name"></a> [named\_value\_name](#output\_named\_value\_name) | The name of the API Management Named Value |
| <a name="output_resource_group_name"></a> [resource\_group\_name](#output\_resource\_group\_name) | The name of the Resource Group |
| <a name="output_secret_named_value_name"></a> [secret\_named\_value\_name](#output\_secret\_named\_value\_name) | The name of the API Management Secret Named Value |
<!-- END_TF_DOCS -->
