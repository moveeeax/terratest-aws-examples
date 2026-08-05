output "vpc_id" {
  description = "The VPC id."
  value       = aws_vpc.this.id
}

output "subnet_ids" {
  description = "Ids of the public subnets, one per AZ."
  value       = aws_subnet.public[*].id
}
