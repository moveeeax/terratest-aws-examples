//go:build e2e

package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/moveeeax/terratest-aws-examples/pkg/awstest"
	"github.com/stretchr/testify/assert"
)

// TestS3Bucket is the canonical plan+apply+assert+destroy lifecycle.
//
// The defer'd Destroy is registered BEFORE Apply so the bucket is torn down
// even if Apply half-succeeds or an assertion below panics — the single most
// important habit for keeping a sandbox account (and its bill) clean.
func TestS3Bucket(t *testing.T) {
	t.Parallel()

	region := awstest.Region()
	name := awstest.UniqueName("tt-s3", random.UniqueId(), 63)

	opts := awstest.Options("../examples/s3-bucket", map[string]interface{}{
		"name":   name,
		"region": region,
		"tags":   awstest.DefaultTags("s3-bucket"),
	})

	defer terraform.Destroy(t, opts)
	terraform.InitAndApply(t, opts)

	bucketID := terraform.Output(t, opts, "bucket_id")
	assert.Equal(t, name, bucketID)

	// The provider guarantees this bucket, but reading versioning back proves
	// the resource is real and configured as the module intends.
	versioning := aws.GetS3BucketVersioning(t, region, bucketID)
	assert.Equal(t, "Enabled", versioning)
}
