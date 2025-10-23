resource "azurerm_api_management_named_value" "named_value" {
  api_management_name = var.api_management_name
  resource_group_name = var.resource_group_name

  name         = var.name
  display_name = var.display_name
  value        = var.value
  secret       = var.secret

  dynamic "value_from_key_vault" {
    for_each = var.value_from_key_vault != null ? ["value_from_key_vault"] : []
    content {
      secret_id          = var.value_from_key_vault.secret_id
      identity_client_id = var.value_from_key_vault.identity_client_id
    }
  }

  tags = var.tags
}
