package cmdparser

type CmdLineArg struct {
	name    string
	options []CmdLineArgOption
}

func (arg *CmdLineArg) Name() string {
	return arg.name
}
