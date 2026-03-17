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

output "id" {
  description = "The ID of the archive (same as the name)."
  value       = aws_cloudwatch_event_archive.archive.id
}

output "arn" {
  description = "The ARN of the archive."
  value       = aws_cloudwatch_event_archive.archive.arn
}

output "name" {
  description = "The name of the archive."
  value       = aws_cloudwatch_event_archive.archive.name
}

output "event_source_arn" {
  description = "The ARN of the event bus associated with the archive."
  value       = aws_cloudwatch_event_archive.archive.event_source_arn
}

output "description" {
  description = "The description of the archive."
  value       = aws_cloudwatch_event_archive.archive.description != null ? aws_cloudwatch_event_archive.archive.description : ""
}

output "event_pattern" {
  description = "The event pattern used to filter events sent to the archive."
  value       = aws_cloudwatch_event_archive.archive.event_pattern != null ? aws_cloudwatch_event_archive.archive.event_pattern : ""
}

output "retention_days" {
  description = "The number of days events are retained in the archive."
  value       = aws_cloudwatch_event_archive.archive.retention_days
}
