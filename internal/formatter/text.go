package formatter

import (
	"fmt"

	"github.com/Madh93/tpm/internal/terraform"
)

type TextFormatter struct{}

func (f *TextFormatter) Format(providers []*terraform.Provider) ([]byte, error) {
	var output string

	if providers != nil {
		for _, provider := range providers {
			output += fmt.Sprintf("%s\n", provider.String())
		}
	} else {
		output = "No packages found.\n"
	}

	return []byte(output), nil
}
