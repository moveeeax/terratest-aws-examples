# terratest-aws-examples

> A copy-paste library of real Terratest patterns for AWS module authors.

**Status:** 🚧 In development

## Overview

Reference collection of Terratest patterns for testing AWS Terraform modules end-to-end.

## Features

- Examples: VPC, S3, IAM, EC2, RDS smoke tests
- Patterns for parallel tests, stages, and retryable AWS eventual consistency
- Fixture layout + `test/` conventions
- CI workflow running tests against a sandbox account with OIDC
- Cost-safe: minimal resources, always destroyed in defer

## Stack

Go + Terratest + Terraform; GitHub Actions with AWS OIDC.

## Usage

```bash
cd test
go test -v -timeout 30m ./...
```

## License

MIT
