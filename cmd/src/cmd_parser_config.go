package main

import (
	cp "cmdparser"
)

// GetCmdParserConfig ...
func GetCmdParserConfig() []cp.CmdLineArg {
	return []cp.CmdLineArg{
		cp.CmdLineArg{
			Name: "init",
			Options: []cp.CmdLineArgOption{
				cp.CmdLineArgOption{
					IsRequired: false,
					Name:       "verbose",
					Params:     []string{}}}},
		cp.CmdLineArg{
			Name: "status",
			Options: []cp.CmdLineArgOption{
				cp.CmdLineArgOption{
					IsRequired: false,
					Name:       "verbose",
					Params:     []string{}}}}}
}
