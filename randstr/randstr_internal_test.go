package randstr

import (
	"strings"
	"testing"
)

func TestGenerateCharset(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		args         CharRanges
		wantPatterns []string
	}{
		{
			args: CharRanges{
				'a': 'b',
			},
			wantPatterns: []string{"ab"},
		},
		{
			args: CharRanges{
				'a': 'b',
				'b': 'c',
			},
			wantPatterns: []string{"ab", "bc"},
		},
		{
			args: CharRanges{
				'a': 'f',
			},
			wantPatterns: []string{"ab", "cdef"},
		},
		{
			args: CharRanges{
				'a': 'z',
			},
			wantPatterns: []string{"abcdefghijklmnopqrstuvwxyz"},
		},
		{
			args: CharRanges{
				'1': '5',
			},
			wantPatterns: []string{"12345"},
		},
		{
			args: CharRanges{
				'1': '5',
				'a': 'd',
			},
			wantPatterns: []string{"abcd", "12345"},
		},
		{
			args: CharRanges{
				'-': '-',
				'_': '_',
			},
			wantPatterns: []string{"-", "_"},
		},
	}

	for _, tc := range testCases {
		got := generateCharset(tc.args)

		for _, pattern := range tc.wantPatterns {
			if !strings.Contains(got, pattern) {
				t.Fatalf(
					"wanted patterns %v to be part of the returned string %s",
					tc.wantPatterns,
					got,
				)
			}
		}
	}
}
