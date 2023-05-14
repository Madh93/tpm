package main

import (
	"log"

	"github.com/Madh93/tpm/cmd"
)

func main() {
	log.SetFlags(log.Default().Flags())
	cmd.Execute()
}
