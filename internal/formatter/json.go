package formatter

import (
	"encoding/json"

	"github.com/Madh93/tpm/internal/terraform"
)

type JSONFormatter struct{}

func (f *JSONFormatter) Format(providers []*terraform.Provider) ([]byte, error) {
	output, err := json.MarshalIndent(providers, "", "  ")
	if err != nil {
		return nil, err
	}

	return output, nil
}
