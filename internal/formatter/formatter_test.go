package formatter_test

import (
	"testing"

	"github.com/Madh93/tpm/internal/formatter"
	"github.com/stretchr/testify/assert"
)

func TestNewFormatter(t *testing.T) {
	tests := []struct {
		name         string
		outputFormat string
		wantErr      bool
	}{
		{
			name:         "valid 'text' output format",
			outputFormat: "text",
			wantErr:      false,
		},
		{
			name:         "valid 'csv' output format",
			outputFormat: "csv",
			wantErr:      false,
		},
		{
			name:         "valid 'json' output format",
			outputFormat: "json",
			wantErr:      false,
		},
		{
			name:         "valid 'table' output format",
			outputFormat: "table",
			wantErr:      false,
		},
		{
			name:         "invalid output format",
			outputFormat: "unknown",
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := formatter.NewFormatter(tt.outputFormat)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, output)
			} else {
				assert.NoError(t, err)
				assert.NotZero(t, output)
			}
		})
	}
}
