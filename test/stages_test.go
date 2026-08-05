//go:build e2e

package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/moveeeax/terratest-aws-examples/pkg/awstest"
	"github.com/stretchr/testify/assert"
)

// TestS3BucketStaged shows terratest's staged pattern. Each stage
// (deploy/validate/teardown) can be skipped independently via env vars, e.g.:
//
//	SKIP_teardown=true go test -run TestS3BucketStaged -tags e2e ./...   # keep infra up
//	SKIP_deploy=true   go test -run TestS3BucketStaged -tags e2e ./...   # re-validate it
//
// That tight edit/apply/inspect loop is the reason to reach for stages instead
// of the one-shot lifecycle in s3_bucket_test.go.
func TestS3BucketStaged(t *testing.T) {
	t.Parallel()

	dir := "../examples/s3-bucket"

	defer test_structure.RunTestStage(t, "teardown", func() {
		opts := test_structure.LoadTerraformOptions(t, dir)
		terraform.Destroy(t, opts)
	})

	test_structure.RunTestStage(t, "deploy", func() {
		opts := awstest.Options(dir, map[string]interface{}{
			"name":   awstest.UniqueName("tt-s3-staged", random.UniqueId(), 63),
			"region": awstest.Region(),
			"tags":   awstest.DefaultTags("s3-bucket"),
		})
		test_structure.SaveTerraformOptions(t, dir, opts)
		terraform.InitAndApply(t, opts)
	})

	test_structure.RunTestStage(t, "validate", func() {
		opts := test_structure.LoadTerraformOptions(t, dir)
		assert.NotEmpty(t, terraform.Output(t, opts, "bucket_arn"))
	})
}
