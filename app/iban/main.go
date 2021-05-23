package main

import (
	"fmt"
	"github.com/yarysh/iban-service/app/iban/cmd"
	"os"
)

type commander interface {
	Run(params []string) error
}

func main() {
	commands := map[string]commander{
		"gen_validator_rules": cmd.GenValidatorRulesCmd{},
	}

	if len(os.Args) < 2 {
		fmt.Println("Error: subcommand required")
		os.Exit(1)
	}

	command, ok := commands[os.Args[1]]
	if !ok {
		fmt.Println("Error: unknown subcommand")
	}

	if err := command.Run(os.Args[2:]); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}
