package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/subratamondal1029/goTodo/internal/database"
	"github.com/subratamondal1029/goTodo/pkgs"
)

func main() {
	args, err := pkgs.GetArguments()

	if err != nil {
		log.Fatal(err)
	}

	switch args.Operation {
	case pkgs.List:
		todos, err := database.Read()

		if err != nil {
			log.Fatal(err)
		}

		if len(todos) == 0 {
			fmt.Println("No todos found.")
			return
		}

		todos.PrintDataTable()
	case pkgs.Add:
		// Add logic for adding a new todo (title)
		title := strings.Join(args.Value, " ")
		todos, err := database.Create(title)

		if err != nil {
			log.Panic(err)
		}

		todos.PrintDataTable()
	case pkgs.ToggleDone:
		// Mark completed (id)
		id := args.Value[0]
		intId, err := strconv.Atoi(id)

		if err != nil {
			log.Panicf("Invalid ID: %s", id)
		}

		todos, err := database.ToggleDone(intId)
		if err != nil {
			log.Panic(err)
		}

		todos.PrintDataTable()
	case pkgs.Delete:
		// Delete item (id)
		id := args.Value[0]
		intId, err := strconv.Atoi(id)

		if err != nil {
			log.Panicf("Invalid ID: %s", id)
		}

		todos, err := database.Delete(intId)
		if err != nil {
			log.Panic(err)
		}

		todos.PrintDataTable()
	}
}
