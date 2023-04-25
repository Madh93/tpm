package parser_test

import (
	"testing"

	"github.com/Madh93/tpm/internal/parser"
	"github.com/stretchr/testify/assert"
)

func TestNewParser(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "valid 'yml' output format",
			path:    "providers.yml",
			wantErr: false,
		},
		{
			name:    "valid 'yaml' output format",
			path:    "whatever.yaml",
			wantErr: false,
		},
		{
			name:    "invalid input format",
			path:    "providers.json",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := parser.NewParser(tt.path)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, p)
			} else {
				assert.NoError(t, err)
				assert.NotZero(t, p)
			}
		})
	}
}
