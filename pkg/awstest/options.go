package awstest

import (
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

// commonRetryableErrors are transient AWS/Terraform failures that are safe to
// retry: eventual-consistency races (a resource that "does not exist" yet) and
// API throttling. Baking them into every example's options is the difference
// between a flaky suite and a reliable one.
var commonRetryableErrors = map[string]string{
	".*does not exist.*":                    "Resource not visible yet (eventual consistency); retrying.",
	".*ResourceNotFoundException.*":         "AWS resource not found yet (eventual consistency); retrying.",
	".*RequestError: send request failed.*": "Transient network error talking to AWS; retrying.",
	".*Throttling.*":                        "AWS API throttling; retrying.",
	".*timeout while waiting for state.*":   "Terraform state wait timed out; retrying.",
}

// Options builds a terraform.Options for an example fixture directory with the
// repository defaults: the caller's vars, the shared retryable-error set, three
// retries with a fixed pause, and colourless output for readable CI logs. The
// returned value is passed straight to terraform.InitAndApply / Destroy.
func Options(dir string, vars map[string]interface{}) *terraform.Options {
	return &terraform.Options{
		TerraformDir:             dir,
		Vars:                     vars,
		RetryableTerraformErrors: commonRetryableErrors,
		MaxRetries:               3,
		TimeBetweenRetries:       5 * time.Second,
		NoColor:                  true,
	}
}
