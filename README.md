# terratest-aws-examples

[![ci](https://github.com/moveeeax/terratest-aws-examples/actions/workflows/ci.yml/badge.svg)](https://github.com/moveeeax/terratest-aws-examples/actions/workflows/ci.yml)
[![Go](https://img.shields.io/badge/Go-1.22-00ADD8?logo=go&logoColor=white)](go.mod)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

A copy-paste library of **real** [Terratest](https://terratest.gruntwork.io/)
patterns for testing AWS Terraform modules end-to-end: the plan/apply/assert/destroy
lifecycle, list-output cross-checks against the live AWS API, retryable
eventual-consistency assertions, and staged tests you can iterate on.

Everything that touches AWS lives under [`test/`](test) behind a build tag, so
you can read the patterns and run the offline unit tests without an AWS account.

## What's here

- **[`pkg/awstest`](pkg/awstest)** — the shared helpers every example uses:
  safe resource naming (`UniqueName`), a standard tag set (`DefaultTags` /
  `MergeTags`), region resolution (`Region`), and a preconfigured
  `terraform.Options` builder with retryable AWS errors baked in. Fully
  unit-tested, no cloud required.
- **[`examples/`](examples)** — three minimal, cost-safe Terraform fixtures:
  `s3-bucket` (versioned, public access blocked), `vpc` (a VPC + public subnets
  per AZ), and `iam-role` (an EC2-assumable role).
- **[`test/`](test)** — the Terratest examples, one per pattern. See
  [`test/README.md`](test/README.md) for the pattern-to-file map.

## How it works

`go test ./...` compiles and runs only the pure-Go helper tests in
`pkg/awstest`. The AWS tests are tagged `//go:build e2e`, so they are invisible
to the default build and only compile/run when you pass `-tags e2e`. CI takes
advantage of that split:

- an **always-on** job builds, vets (including the `e2e` tag so the AWS tests
  can never rot), runs the unit tests with the race detector, and
  `tofu validate`s every example — all offline;
- an **opt-in** job (`workflow_dispatch`) assumes a short-lived AWS role via
  OIDC and actually applies the examples against a sandbox account. No
  long-lived keys, and it never runs on a plain push.

## Usage

Run the offline suite exactly as CI does:

```console
$ go test ./...
ok   github.com/moveeeax/terratest-aws-examples/pkg/awstest   0.034s
```

Run the real AWS examples against a sandbox (creates and destroys resources):

```console
$ export AWS_REGION=us-east-1
$ go test -tags e2e -v -timeout 30m ./test/...
=== RUN   TestS3Bucket
--- PASS: TestS3Bucket (23.11s)
=== RUN   TestVpc
--- PASS: TestVpc (41.78s)
...
```

Keep infrastructure up while iterating on a single example:

```console
$ SKIP_teardown=true go test -tags e2e -run TestS3BucketStaged ./test/...
```

## One note

Read outputs back from AWS, not just from Terraform state. `vpc_test.go` asserts
the module's `subnet_ids` output *and* calls `aws.GetSubnetsForVpc` — because a
green `terraform apply` proves the API accepted your request, not that the
resource is really there and shaped the way you think. The extra call is what
turns a smoke test into a real one.

## License

MIT — see [LICENSE](LICENSE).
