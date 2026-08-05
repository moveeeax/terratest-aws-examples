variable "name" {
  description = "IAM role name."
  type        = string
}

variable "region" {
  description = "AWS region for the provider (IAM itself is global)."
  type        = string
  default     = "us-east-1"
}

variable "tags" {
  description = "Tags applied to the role."
  type        = map(string)
  default     = {}
}
