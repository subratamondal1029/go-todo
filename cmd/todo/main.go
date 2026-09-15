package main

import (
	"fmt"
	"log"

	"github.com/subratamondal1029/goTodo/internal/cli"
	"github.com/subratamondal1029/goTodo/internal/database"
)

func main() {
	args, err := cli.GetArguments()

	if err != nil {
		log.Panic(err)
	}

	fmt.Println(args)

	connection, err := database.Connect(cli.DBURI)
	if err != nil {
		log.Panic(err)
	}

	defer connection.DB.Close()

	fmt.Println("Database connected.")
}
