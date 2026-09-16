package todo

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/subratamondal1029/goTodo/internal/cli"
	"github.com/subratamondal1029/goTodo/internal/database"
)

func strToInt(s string) (int, error) {
	intS, err := strconv.Atoi(s)
	intS = (intS)

	if err != nil {
		return 0, err
	}

	return intS, nil
}

func Resolve(args *cli.Argument, connection *database.Connection) error {
	service := NewTodoService(*connection.Queries)

	switch args.Operation {
	case cli.List:
		todos, err := service.ReadAll()
		if err != nil {
			return err
		}

		PrintDataTable(&todos)

	case cli.Get:
		if len(args.Value) < 1 {
			return fmt.Errorf("Arguments underload \n usage: todo %s <id>", cli.Get)
		}

		id := args.Value[0]
		idInt, err := strToInt(id)

		if err != nil {
			return fmt.Errorf("Invalid ID: %s", id)
		}

		todo, err := service.ReadOne(idInt)
		if err != nil {
			return err
		}

		PrintDataTable(&[]database.Todo{*todo})

	case cli.Add:
		if len(args.Value) < 1 {
			return fmt.Errorf("Arguments underload \n usage: todo %s <title>", cli.Add)
		}

		title := strings.Join(args.Value, " ")
		if title == "" {
			return fmt.Errorf("Title cannot be empty")
		}

		todo, err := service.Create(title)

		if err != nil {
			return err
		}

		PrintDataTable(&[]database.Todo{*todo})

	case cli.Done:
		if len(args.Value) < 2 {
			return fmt.Errorf("Arguments underload \n usage: todo %s <id> <true|false>", cli.Done)
		}

		id := args.Value[0]
		done := args.Value[1]

		idInt, err := strToInt(id)

		if err != nil {
			return fmt.Errorf("Invalid ID: %s", id)
		}

		doneBool, err := strconv.ParseBool(done)

		if err != nil {
			return fmt.Errorf("Invalid done value: %s", done)
		}

		todo, err := service.ToggleDone(idInt, doneBool)
		if err != nil {
			return err
		}

		PrintDataTable(&[]database.Todo{*todo})

	case cli.Delete:
		if len(args.Value) < 1 {
			return fmt.Errorf("Arguments underload \n usage: todo %s <id>", cli.Delete)
		}

		id := args.Value[0]
		idInt, err := strToInt(id)

		if err != nil {
			return fmt.Errorf("Invalid ID: %s", id)
		}

		err = service.Delete(idInt)
		if err != nil {
			return err
		}
		fmt.Printf("Deleted todo with ID: %d\n", idInt)
	default:
		return fmt.Errorf("Unknown operation: %s", args.Operation)
	}

	return nil
}
