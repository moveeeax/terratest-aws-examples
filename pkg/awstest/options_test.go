package awstest

import "testing"

func TestOptions(t *testing.T) {
	vars := map[string]interface{}{"name": "tt-example", "region": "us-east-1"}
	opts := Options("../../examples/s3-bucket", vars)

	if opts.TerraformDir != "../../examples/s3-bucket" {
		t.Errorf("TerraformDir = %q", opts.TerraformDir)
	}
	if opts.MaxRetries != 3 {
		t.Errorf("MaxRetries = %d, want 3", opts.MaxRetries)
	}
	if opts.TimeBetweenRetries <= 0 {
		t.Errorf("TimeBetweenRetries not set: %v", opts.TimeBetweenRetries)
	}
	if !opts.NoColor {
		t.Error("NoColor should be true for CI logs")
	}
	if len(opts.RetryableTerraformErrors) == 0 {
		t.Error("RetryableTerraformErrors should be populated")
	}
	if opts.Vars["name"] != "tt-example" {
		t.Errorf("Vars not passed through: %v", opts.Vars)
	}
}
