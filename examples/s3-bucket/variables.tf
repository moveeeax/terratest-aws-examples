variable "name" {
  description = "Globally unique bucket name."
  type        = string
}

variable "region" {
  description = "AWS region to create the bucket in."
  type        = string
  default     = "us-east-1"
}

variable "tags" {
  description = "Tags applied to the bucket."
  type        = map(string)
  default     = {}
}
