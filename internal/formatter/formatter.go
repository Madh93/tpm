package formatter

import (
	"fmt"

	"github.com/Madh93/tpm/internal/terraform"
)

var (
	ProviderHeader = []string{"namespace", "name", "version", "os", "arch"}
)

type OutputFormatter interface {
	Format(providers []*terraform.Provider) ([]byte, error)
}

func NewFormatter(outputFormat string) (OutputFormatter, error) {
	switch outputFormat {
	case "text":
		return &TextFormatter{}, nil
	case "csv":
		return &CSVFormatter{}, nil
	case "json":
		return &JSONFormatter{}, nil
	case "table":
		return &TableFormatter{}, nil
	}
	return nil, fmt.Errorf("unsupported '%s' output format", outputFormat)
}
