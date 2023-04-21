package terraform

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProviderVersionsLast(t *testing.T) {
	tests := []struct {
		name     string
		versions ProviderVersions
		expected *ProviderVersion
	}{
		{
			name:     "valid version",
			versions: ProviderVersions{{Version: "1.0.0"}, {Version: "1.1.0"}, {Version: "1.2.0"}},
			expected: &ProviderVersion{Version: "1.2.0"},
		},
		{
			name:     "another valid version",
			versions: ProviderVersions{{Version: "2.0.0"}, {Version: "1.5.0"}, {Version: "1.7.0"}},
			expected: &ProviderVersion{Version: "1.7.0"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.versions.Last())
		})
	}
}

func TestProviderVersionsSort(t *testing.T) {
	tests := []struct {
		name     string
		versions ProviderVersions
		expected ProviderVersions
	}{
		{
			name:     "valid versions",
			versions: ProviderVersions{{Version: "1.2.0"}, {Version: "1.0.0"}, {Version: "1.1.0"}},
			expected: ProviderVersions{{Version: "1.0.0"}, {Version: "1.1.0"}, {Version: "1.2.0"}},
		},
		{
			name:     "more valid versions",
			versions: ProviderVersions{{Version: "2.4.0"}, {Version: "3.0.1"}, {Version: "1.2.0"}},
			expected: ProviderVersions{{Version: "1.2.0"}, {Version: "2.4.0"}, {Version: "3.0.1"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Sort(tt.versions)
			for i, version := range tt.versions {
				assert.Equal(t, tt.expected[i], version)

			}
		})
	}
}
