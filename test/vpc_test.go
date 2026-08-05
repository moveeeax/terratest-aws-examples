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

// TestVpc shows how to assert on a list output and cross-check it against the
// AWS API rather than trusting Terraform's own state. It queries the subnets
// AWS actually reports for the VPC and compares that count to what the module
// says it created.
func TestVpc(t *testing.T) {
	t.Parallel()

	region := awstest.Region()
	name := awstest.UniqueName("tt-vpc", random.UniqueId(), 40)
	const subnetCount = 2

	opts := awstest.Options("../examples/vpc", map[string]interface{}{
		"name":         name,
		"region":       region,
		"subnet_count": subnetCount,
		"tags":         awstest.DefaultTags("vpc"),
	})

	defer terraform.Destroy(t, opts)
	terraform.InitAndApply(t, opts)

	vpcID := terraform.Output(t, opts, "vpc_id")
	subnetIDs := terraform.OutputList(t, opts, "subnet_ids")
	assert.Len(t, subnetIDs, subnetCount)

	subnets := aws.GetSubnetsForVpc(t, vpcID, region)
	assert.Len(t, subnets, subnetCount, "AWS should report the same subnet count the module output claims")
}
