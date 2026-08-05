//go:build e2e

package test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/retry"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/moveeeax/terratest-aws-examples/pkg/awstest"
	"github.com/stretchr/testify/assert"
)

// TestIamRole demonstrates the retryable eventual-consistency pattern. IAM is
// global and lags behind writes, so instead of asserting once we poll with
// retry.DoWithRetry until the role's ARN settles into the expected shape.
func TestIamRole(t *testing.T) {
	t.Parallel()

	name := awstest.UniqueName("tt-role", random.UniqueId(), 64)

	opts := awstest.Options("../examples/iam-role", map[string]interface{}{
		"name":   name,
		"region": awstest.Region(),
		"tags":   awstest.DefaultTags("iam-role"),
	})

	defer terraform.Destroy(t, opts)
	terraform.InitAndApply(t, opts)

	roleName := terraform.Output(t, opts, "role_name")
	assert.Equal(t, name, roleName)

	// Poll for up to ~30s. Returning a plain error makes DoWithRetry sleep and
	// try again; if the ARN never becomes valid it fails the test after the
	// last attempt. This is the idiomatic guard against eventual consistency.
	retry.DoWithRetry(t, "wait for role arn", 10, 3*time.Second, func() (string, error) {
		arn := terraform.Output(t, opts, "role_arn")
		if !strings.HasPrefix(arn, "arn:aws:iam::") || !strings.HasSuffix(arn, ":role/"+name) {
			return "", fmt.Errorf("role arn not ready or malformed: %q", arn)
		}
		return arn, nil
	})
}
