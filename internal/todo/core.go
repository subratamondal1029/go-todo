package todo

import (
	"fmt"
	"log"
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

func Resolve(args *cli.Argument, connection *database.Connection) {
	service := NewTodoService(*connection.Queries)

	switch args.Operation {
	case cli.List:
		todos, err := service.ReadAll()
		if err != nil {
			log.Fatal(err)
		}

		PrintDataTable(&todos)

	case cli.Get:
		id := args.Value[0]
		idInt, err := strToInt(id)

		if err != nil {
			log.Fatalf("Invalid ID: %s", id)
		}

		todo, err := service.ReadOne(idInt)
		if err != nil {
			log.Fatal(err)
		}

		PrintDataTable(&[]database.Todo{*todo})

	case cli.Add:
		title := strings.Join(args.Value, " ")
		if title == "" {
			log.Fatal("Title cannot be empty")
		}

		todo, err := service.Create(title)

		if err != nil {
			log.Fatal(err)
		}

		PrintDataTable(&[]database.Todo{*todo})

	case cli.ToggleDone:
		id := args.Value[0]
		done := args.Value[1]

		idInt, err := strToInt(id)

		if err != nil {
			log.Fatalf("Invalid ID: %s", id)
		}

		doneBool, err := strconv.ParseBool(done)

		if err != nil {
			log.Fatalf("Invalid done value: %s", done)
		}

		todo, err := service.ToggleDone(idInt, doneBool)
		if err != nil {
			log.Fatal(err)
		}

		PrintDataTable(&[]database.Todo{*todo})

	case cli.Delete:
		id := args.Value[0]
		idInt, err := strToInt(id)

		if err != nil {
			log.Fatalf("Invalid ID: %s", id)
		}

		err = service.Delete(idInt)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Deleted todo with ID: %d\n", idInt)
	}
}
