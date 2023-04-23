package formatter

import (
	"bytes"
	"encoding/csv"

	"github.com/Madh93/tpm/internal/terraform"
)

type CSVFormatter struct{}

func (f *CSVFormatter) Format(providers []*terraform.Provider) ([]byte, error) {
	var output bytes.Buffer
	writer := csv.NewWriter(&output)

	err := writer.Write(ProviderHeader)
	if err != nil {
		return nil, err
	}

	for _, provider := range providers {
		err := writer.Write(provider.ToOutputRow())
		if err != nil {
			return nil, err
		}
	}

	writer.Flush()
	return output.Bytes(), nil
}
