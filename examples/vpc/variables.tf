variable "name" {
  description = "Name prefix for the VPC and its subnets."
  type        = string
}

variable "region" {
  description = "AWS region to create the VPC in."
  type        = string
  default     = "us-east-1"
}

variable "cidr_block" {
  description = "VPC CIDR block. Subnets are carved from this with a /8 offset."
  type        = string
  default     = "10.0.0.0/16"
}

variable "subnet_count" {
  description = "Number of public subnets to create, one per AZ."
  type        = number
  default     = 2
}

variable "tags" {
  description = "Tags applied to the VPC and subnets."
  type        = map(string)
  default     = {}
}
