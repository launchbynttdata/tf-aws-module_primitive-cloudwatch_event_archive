// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

# -----------------------------------------------------------------------------
# Required
# -----------------------------------------------------------------------------

variable "name" {
  description = "Name of the EventBridge event archive. Must be unique per account and region. Maximum 48 characters."
  type        = string

  validation {
    condition     = length(var.name) >= 1 && length(var.name) <= 48
    error_message = "Archive name must be between 1 and 48 characters."
  }
}

variable "event_source_arn" {
  description = "ARN of the event bus associated with the archive. Events from this bus are archived."
  type        = string
}

# -----------------------------------------------------------------------------
# Optional
# -----------------------------------------------------------------------------

variable "description" {
  description = "Description for the archive. Maximum 512 characters."
  type        = string
  default     = null

  validation {
    condition     = var.description == null ? true : (length(var.description) >= 0 && length(var.description) <= 512)
    error_message = "Description must be between 0 and 512 characters."
  }
}

variable "event_pattern" {
  description = "Event pattern to filter events sent to the archive. JSON string. Archives all events if not specified."
  type        = string
  default     = null
}

variable "retention_days" {
  description = "Maximum number of days to retain events in the archive. Omit for indefinite retention."
  type        = number
  default     = null

  validation {
    condition     = var.retention_days == null ? true : var.retention_days >= 0
    error_message = "Retention days must be 0 or greater."
  }
}

variable "kms_key_identifier" {
  description = "The ARN, Key ID, or alias of the KMS key EventBridge uses to encrypt the archive. Omit to use the AWS owned key."
  type        = string
  default     = null
}
