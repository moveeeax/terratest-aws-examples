package awstest

import "testing"

func TestUniqueName(t *testing.T) {
	cases := []struct {
		name           string
		prefix, suffix string
		max            int
		want           string
	}{
		{"simple", "tt-s3", "AbC12", 63, "tt-s3-abc12"},
		{"uppercase collapsed", "TT_S3", "Bucket", 63, "tt-s3-bucket"},
		{"symbols collapsed", "tt..s3", "a__b", 63, "tt-s3-a-b"},
		{"empty suffix", "tt-vpc", "", 63, "tt-vpc"},
		{"empty prefix", "", "role", 63, "role"},
		{"truncated no trailing dash", "tt-bucket", "0123456789", 12, "tt-bucket-01"},
		{"truncation trims dash", "tt-bucket", "xxxxxxxxxx", 10, "tt-bucket"},
		{"no limit", "tt", "long-suffix", 0, "tt-long-suffix"},
		{"leading trailing symbols trimmed", "-tt-", "-s3-", 63, "tt-s3"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := UniqueName(tc.prefix, tc.suffix, tc.max)
			if got != tc.want {
				t.Fatalf("UniqueName(%q,%q,%d) = %q, want %q", tc.prefix, tc.suffix, tc.max, got, tc.want)
			}
			if tc.max > 0 && len(got) > tc.max {
				t.Fatalf("result %q exceeds max %d", got, tc.max)
			}
		})
	}
}
