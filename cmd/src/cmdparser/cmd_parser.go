package cmdparser

import (
	"fmt"
	"os"
)

type CmdParser struct {
	parsedArgs []string     // Args as strings that were passed by the user
	validArgs  []CmdLineArg // Args that are valid
	currentArg []CmdLineArg // Arg that was successfully parsed
}

// Init ...
func (parser *CmdParser) Init(args []CmdLineArg) {
	parser.validArgs = args
}

func (parser *CmdParser) ValidateCmdLineArgs() bool {

	result := false

	if len(parser.parsedArgs) > 1 {

		// First arg is path on windows
		cmdArg := parser.parsedArgs[1]

		// Search in list of valid args is given arg is correct
		for i := 0; i < len(parser.validArgs); i++ {

			currentValidArg := parser.validArgs[i]

			if currentValidArg.Name() == cmdArg {
				result = true
				break
			}
		}
	}

	return result
}

func (parser *CmdParser) checkArgsSyntax() bool {
	return true
}

func (parser *CmdParser) SaveCmdLineArgs() {
	args := os.Args

	// Save all args in member list
	for _, currentArg := range args {
		parser.parsedArgs = append(parser.parsedArgs, currentArg)
	}
}

// test
func (parser *CmdParser) PrintCmdLineArgs() {

	for _, currentArg := range parser.parsedArgs {
		fmt.Printf("%s\n", currentArg)
	}
}
