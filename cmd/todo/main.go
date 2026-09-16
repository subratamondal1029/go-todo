package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/subratamondal1029/goTodo/internal/cli"
	"github.com/subratamondal1029/goTodo/internal/database"
	"github.com/subratamondal1029/goTodo/internal/todo"
)

func main() {
	args, err := cli.GetArguments()

	if err != nil {
		log.Panic(err)
	}

	connection, err := database.Connect(cli.DBURI)
	if err != nil {
		log.Panic(err)
	}

	defer connection.DB.Close()

	fmt.Println("Database connected.")
	fmt.Println(strings.Repeat("=", 20))
	todo.Resolve(args, connection)
}
