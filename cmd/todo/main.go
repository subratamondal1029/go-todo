package main

import (
	"fmt"
	"log"

	"github.com/subratamondal1029/goTodo/internal/cli"
)

func main() {
	args, err := cli.GetArguments()

	if err != nil {
		log.Panic(err)
	}

	fmt.Println(args)
	fmt.Printf("Db uri: %s\n", cli.DBURI)
}
