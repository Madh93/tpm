package parser

import (
	"fmt"
	"path/filepath"

	"github.com/Madh93/tpm/internal/terraform"
)

type InputParser interface {
	Parse(input []byte) ([]*terraform.Provider, error)
}

func NewParser(path string) (InputParser, error) {
	extension := filepath.Ext(path)

	switch extension {
	case ".yml", ".yaml":
		return &YAMLParser{}, nil
	}

	return nil, fmt.Errorf("unsupported '%s' input format", extension)
}
