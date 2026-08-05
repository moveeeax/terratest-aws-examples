// Package awstest provides small, dependency-light helpers that the Terratest
// examples in this repository share: safe resource naming, a standard tag set,
// region resolution, and a preconfigured terraform.Options builder. Keeping
// these in one tested package means every example names, tags, and retries the
// same way, so a sweeper can always find and destroy leaked resources.
package awstest

import (
	"regexp"
	"strings"
)

var nonAlnum = regexp.MustCompile(`[^a-z0-9]+`)

// UniqueName builds an AWS- and DNS-safe resource name from prefix and suffix.
//
// The result is lowercased, every run of non-alphanumeric characters is
// collapsed to a single '-', leading/trailing '-' are trimmed, and the whole
// string is truncated to max characters (max <= 0 means no limit). Truncation
// never leaves a trailing '-', so the name stays valid even after clipping.
//
// Typical use pairs a static prefix with terratest's random.UniqueId():
//
//	name := UniqueName("tt-s3", random.UniqueId(), 63)
func UniqueName(prefix, suffix string, max int) string {
	parts := make([]string, 0, 2)
	for _, p := range []string{prefix, suffix} {
		p = nonAlnum.ReplaceAllString(strings.ToLower(p), "-")
		p = strings.Trim(p, "-")
		if p != "" {
			parts = append(parts, p)
		}
	}
	name := strings.Join(parts, "-")
	if max > 0 && len(name) > max {
		name = strings.TrimRight(name[:max], "-")
	}
	return name
}
