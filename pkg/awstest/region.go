package awstest

import "os"

// DefaultRegion is used when no region is set in the environment.
const DefaultRegion = "us-east-1"

// Region resolves the AWS region for a test. It honours TERRATEST_REGION first
// (so CI can pin the examples to a sandbox region without touching a
// developer's AWS_REGION), then AWS_REGION, then AWS_DEFAULT_REGION, and
// finally falls back to DefaultRegion.
func Region() string {
	for _, key := range []string{"TERRATEST_REGION", "AWS_REGION", "AWS_DEFAULT_REGION"} {
		if v := os.Getenv(key); v != "" {
			return v
		}
	}
	return DefaultRegion
}
