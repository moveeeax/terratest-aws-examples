package awstest

import "testing"

func TestRegionDefault(t *testing.T) {
	t.Setenv("TERRATEST_REGION", "")
	t.Setenv("AWS_REGION", "")
	t.Setenv("AWS_DEFAULT_REGION", "")
	if got := Region(); got != DefaultRegion {
		t.Errorf("Region() = %q, want %q", got, DefaultRegion)
	}
}

func TestRegionPrecedence(t *testing.T) {
	t.Setenv("AWS_DEFAULT_REGION", "us-west-2")
	t.Setenv("AWS_REGION", "eu-west-1")
	if got := Region(); got != "eu-west-1" {
		t.Errorf("AWS_REGION should win over AWS_DEFAULT_REGION: got %q", got)
	}

	t.Setenv("TERRATEST_REGION", "ap-south-1")
	if got := Region(); got != "ap-south-1" {
		t.Errorf("TERRATEST_REGION should win: got %q", got)
	}
}
