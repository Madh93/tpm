package mathutils_test

import (
	"testing"

	"github.com/Madh93/tpm/internal/mathutils"
	"github.com/stretchr/testify/assert"
)

func TestMaxWithInt(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{
			a:        1,
			b:        2,
			expected: 2,
		},
	}

	for _, tc := range tests {
		got := mathutils.Max(tc.a, tc.b)
		assert.Equal(t, tc.expected, got)
	}
}

func TestMaxWithFloat(t *testing.T) {
	tests := []struct {
		a, b, expected float32
	}{
		{
			a:        5.5,
			b:        -2.4,
			expected: 5.5,
		},
	}

	for _, tc := range tests {
		got := mathutils.Max(tc.a, tc.b)
		assert.Equal(t, tc.expected, got)
	}
}

func TestMaxWithString(t *testing.T) {
	tests := []struct {
		a, b, expected string
	}{
		{
			a:        "foo",
			b:        "bar",
			expected: "foo",
		},
	}

	for _, tc := range tests {
		got := mathutils.Max(tc.a, tc.b)
		assert.Equal(t, tc.expected, got)
	}
}

func TestMinWithInt(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{
			a:        1,
			b:        2,
			expected: 1,
		},
	}

	for _, tc := range tests {
		got := mathutils.Min(tc.a, tc.b)
		assert.Equal(t, tc.expected, got)
	}
}

func TestMinWithFloat(t *testing.T) {
	tests := []struct {
		a, b, expected float32
	}{
		{
			a:        5.5,
			b:        -2.4,
			expected: -2.4,
		},
	}

	for _, tc := range tests {
		got := mathutils.Min(tc.a, tc.b)
		assert.Equal(t, tc.expected, got)
	}
}

func TestMinWithString(t *testing.T) {
	tests := []struct {
		a, b, expected string
	}{
		{
			a:        "foo",
			b:        "bar",
			expected: "bar",
		},
	}

	for _, tc := range tests {
		got := mathutils.Min(tc.a, tc.b)
		assert.Equal(t, tc.expected, got)
	}
}
