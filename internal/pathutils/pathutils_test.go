package pathutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathDepth(t *testing.T) {
	tests := []struct {
		path     string
		expected int
	}{
		{
			path:     "",
			expected: 0,
		},
		{
			path:     "/",
			expected: 0,
		},
		{
			path:     "/a",
			expected: 1,
		},
		{
			path:     "/a/b",
			expected: 2,
		},
		{
			path:     "/a/b/c",
			expected: 3,
		},
		{
			path:     "/a/b/c/d",
			expected: 4,
		},
	}

	for _, tc := range tests {
		got := PathDepth(tc.path)
		assert.Equal(t, tc.expected, got)
	}
}
