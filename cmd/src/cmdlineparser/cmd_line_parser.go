package cmdlineparser

import (
	"fmt"
	"os"
)

// test
func PrintCmdArgs() {

	args := os.Args

	for _, currentArg := range args {
		fmt.Printf("%s\n", currentArg)
	}
}
