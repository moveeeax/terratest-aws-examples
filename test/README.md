# test/ — Terratest examples

Every file here is tagged `//go:build e2e`, so `go test ./...` (the default) does
**not** compile or run them. That is deliberate: these tests call
`terraform apply` against a real AWS account and cost money. The offline unit
tests for the shared helpers live under [`pkg/awstest`](../pkg/awstest) and run
on every push.

## Run the examples

You need AWS credentials with permission to create/destroy S3, EC2 (VPC), and
IAM resources, plus `terraform` (or OpenTofu symlinked as `terraform`) on PATH.

```bash
export AWS_REGION=us-east-1   # or TERRATEST_REGION to pin only the tests
go test -tags e2e -v -timeout 30m ./test/...
```

## The patterns, and when to reach for each

| File | Pattern | Use it when |
|------|---------|-------------|
| `s3_bucket_test.go` | plan → apply → assert → **defer destroy** | The default lifecycle for any single module. |
| `vpc_test.go` | assert a list output, then cross-check it against the live AWS API | You do not trust Terraform state alone and want AWS to confirm. |
| `iam_role_test.go` | `retry.DoWithRetry` around an assertion | The resource is eventually consistent (IAM, DNS, ACM). |
| `stages_test.go` | `test_structure` deploy/validate/teardown stages | You are iterating and want to keep infra up between runs (`SKIP_teardown=true`). |

All examples name resources with `awstest.UniqueName(...)` and tag them with
`awstest.DefaultTags(...)`, so a leaked resource always carries
`Project=terratest-aws-examples` and can be swept.
