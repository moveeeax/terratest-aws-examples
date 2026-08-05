package awstest

import "testing"

func TestDefaultTags(t *testing.T) {
	tags := DefaultTags("s3-bucket")
	if tags["Project"] != Project {
		t.Errorf("Project = %q, want %q", tags["Project"], Project)
	}
	if tags["Module"] != "s3-bucket" {
		t.Errorf("Module = %q, want s3-bucket", tags["Module"])
	}
	if tags["Test"] != "true" {
		t.Errorf("Test = %q, want true", tags["Test"])
	}
}

func TestMergeTags(t *testing.T) {
	base := DefaultTags("vpc")
	got := MergeTags(base, map[string]string{"Owner": "ci", "Module": "override"})

	if got["Owner"] != "ci" {
		t.Errorf("Owner = %q, want ci", got["Owner"])
	}
	if got["Module"] != "override" {
		t.Errorf("override should win: Module = %q, want override", got["Module"])
	}
	// base must not be mutated.
	if base["Module"] != "vpc" {
		t.Errorf("base was mutated: Module = %q, want vpc", base["Module"])
	}
	if base["Owner"] != "" {
		t.Errorf("base leaked Owner tag: %q", base["Owner"])
	}
}

func TestMergeTagsNil(t *testing.T) {
	got := MergeTags(nil, nil)
	if len(got) != 0 {
		t.Errorf("MergeTags(nil,nil) = %v, want empty", got)
	}
}
