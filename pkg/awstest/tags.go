package awstest

// Project is the tag value stamped on every resource the examples create. A
// crashed test can leak resources; tagging them all with Project=<this> lets a
// sweeper list and delete anything left behind, regardless of which example
// created it.
const Project = "terratest-aws-examples"

// DefaultTags returns the baseline tags applied to every example resource.
// module identifies which example owns the resource (e.g. "s3-bucket").
func DefaultTags(module string) map[string]string {
	return map[string]string{
		"Project":   Project,
		"Module":    module,
		"ManagedBy": "terraform",
		"Test":      "true",
	}
}

// MergeTags returns a new map with base overlaid by override. override wins on
// key conflicts; neither input is mutated and nil inputs are treated as empty.
func MergeTags(base, override map[string]string) map[string]string {
	out := make(map[string]string, len(base)+len(override))
	for k, v := range base {
		out[k] = v
	}
	for k, v := range override {
		out[k] = v
	}
	return out
}
