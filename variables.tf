variable "resource_group_name" {
  type        = string
  description = "name of the resource group where the APIM exists"
  default     = null
  validation {
    condition     = can(regex("^[a-zA-Z0-9-]{1,50}$", var.resource_group_name))
    error_message = "The resource group name can only contain alphanumeric characters and dashes and must be between 1 and 50 characters long."
  }
}

variable "api_management_name" {
  type        = string
  description = "name of the APIM in which this named value will de deployed"
  default     = null
  validation {
    condition     = can(regex("^[a-zA-Z0-9-]{1,50}$", var.api_management_name))
    error_message = "The APIM name can only contain alphanumeric characters and dashes and must be between 1 and 50 characters long."
  }
}

variable "name" {
  type        = string
  description = "The name of the named value."
}

variable "display_name" {
  type        = string
  description = "The display name of the named value."
}

variable "value" {
  type        = string
  description = "The value of the named value."
  default     = null
}

variable "value_from_key_vault" {
  type = object({
    secret_id          = string
    identity_client_id = optional(string, null)
  })
  description = "The Key Vault secret identifier to reference for the named value's value."
  default     = null
}

variable "secret" {
  type        = bool
  description = "Specifies whether the named value is a secret."
  default     = false
}

variable "tags" {
  type        = set(string)
  description = "A set of tags to assign to the resource."
  default     = []
}
