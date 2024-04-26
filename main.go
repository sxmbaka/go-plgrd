package main

import (
	"os"

	"github.com/sxmbaka/go-plgrd/general"
	"github.com/sxmbaka/go-plgrd/playground"
	"github.com/sxmbaka/go-plgrd/variables"
)

func main() {
	// Get command-line arguments excluding the program name
	args := os.Args[1:]
	if len(args) > 0 {
		switch args[0] {
		case "variables", "var":
			variables.TestAllOperations()
			variables.TestAllVariables()
			variables.TypeConversion()
			variables.TypeInference()
			variables.TestAllScopes()
			variables.ShadowingVariables()
		case "constants", "const":
			variables.TestConstants()
		case "iota":
			general.TestIota()
		case "playground", "p":
			playground.Test()
		default:
			// Handle invalid command-line argument
			println("Invalid argument")
		}
	} else {
		// Handle no command-line argument to run all
	}
}
