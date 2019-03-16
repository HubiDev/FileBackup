package cmdparser

import (
	"fmt"
	"os"
)

type CmdLineArg struct {
	argName    string
	argOptions []CmdLineArgOption
}

// CmdLineArgOption ...
type CmdLineArgOption struct {
	isRequired   bool
	optionName   string
	optionParams []string
}

type CmdParser struct {

	// Args as strings that were passed by the user
	parsedArgs []string

	// Args that are known as valid by the parser
	args []CmdLineArg
}

func (parser *CmdParser) SaveCmdLineArgs() {
	args := os.Args

	// Save all args in member list
	for _, currentArg := range args {
		parser.parsedArgs = append(parser.parsedArgs, currentArg)
	}
}

func (parser *CmdParser) ValidateCmdLineArgs() bool {

	return true
}

func (parser *CmdParser) checkArgsSyntax() bool {
	return true
}

// test
func (parser *CmdParser) PrintCmdLineArgs() {

	for _, currentArg := range parser.parsedArgs {
		fmt.Printf("%s\n", currentArg)
	}
}
