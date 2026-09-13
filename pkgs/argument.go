package pkgs

import (
	"fmt"
	"os"
)

type Operation string

const (
	Add    Operation = "add"
	List   Operation = "list"
	Done   Operation = "done"
	Delete Operation = "delete"
)

type Argument struct {
	Operation Operation
	Value     []string
}

func GetArguments() (*Argument, error) {
	args := os.Args[1:]

	if len(args) == 0 {
		return nil, fmt.Errorf("no arguments provided")
	}

	var arg Argument

	switch args[0] {
	case string(Add):
		arg = Argument{Operation: Add}
	case string(List):
		arg = Argument{Operation: List}
	case string(Done):
		arg = Argument{Operation: Done}
	case string(Delete):
		arg = Argument{Operation: Delete}
	default:
		return nil, fmt.Errorf("invalid operation: %s", args[0])
	}

	arg.Value = args[1:]

	return &arg, nil
}
