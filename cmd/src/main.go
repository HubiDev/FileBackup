package main

import (
	"cmdparser"
	"fmt"
)

func main() {

	fmt.Print("FileBackup version 0.0.1\n")

	cmdParser := cmdparser.CmdParser{}
	cmdParser.Init(GetCmdParserConfig())
	cmdParser.SaveCmdLineArgs()
	cmdParser.PrintCmdLineArgs()
}
