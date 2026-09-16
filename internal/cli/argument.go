package cli

import (
	"fmt"
	"os"
)

type Operation string

const (
	Add    Operation = "add"
	List   Operation = "list"
	Get    Operation = "get"
	Done   Operation = "done"
	Delete Operation = "delete"
)

type Argument struct {
	Skip      bool
	Operation Operation
	Value     []string
}

func GetArguments() (*Argument, error) {
	args := os.Args[1:]

	if len(args) == 0 {
		return &Argument{Skip: true}, nil
	}

	var arg Argument

	switch args[0] {
	case string(Add):
		arg = Argument{Operation: Add}
	case string(Get):
		arg = Argument{Operation: Get}
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
