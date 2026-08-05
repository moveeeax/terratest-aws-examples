output "role_name" {
  description = "The IAM role name."
  value       = aws_iam_role.this.name
}

output "role_arn" {
  description = "The IAM role ARN."
  value       = aws_iam_role.this.arn
}
