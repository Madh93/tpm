package formatter

import (
	"bytes"

	"github.com/Madh93/tpm/internal/terraform"
	"github.com/olekukonko/tablewriter"
)

type TableFormatter struct{}

func (f *TableFormatter) Format(providers []*terraform.Provider) ([]byte, error) {
	var output bytes.Buffer
	table := tablewriter.NewWriter(&output)

	table.SetHeader(ProviderHeader)

	for _, provider := range providers {
		table.Append(provider.ToOutputRow())
	}

	table.Render()
	return output.Bytes(), nil
}
