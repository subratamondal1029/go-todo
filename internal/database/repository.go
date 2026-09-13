package database

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/subratamondal1029/goTodo/pkgs"
)

type Todos []Todo

func checkExistingTodo(todos Todos, title string) bool {
	for _, todo := range todos {
		if todo.Title == title {
			return true
		}
	}

	return false
}

func (todos Todos) PrintDataTable() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer w.Flush()

	fmt.Fprintln(w, "ID\tTITLE\tCOMPLETED")
	for _, todo := range todos {
		fmt.Fprintf(w, "%d\t%s\t%t\n",
			todo.Id,
			todo.Title,
			todo.Completed,
		)
	}
}

func Read() (Todos, error) {
	readFile, err := os.Open(pkgs.AppStateFile)

	if err != nil {
		return nil, fmt.Errorf("Error Reading state file: %w", err)
	}
	defer readFile.Close()

	var todos Todos
	decoder := json.NewDecoder(readFile)

	if err = decoder.Decode(&todos); err != nil {
		return nil, fmt.Errorf("Error Decoding state file: %w", err)
	}

	return todos, nil
}

func Create(title string) (Todos, error) {
	todos, err := Read()

	if err != nil {
		return nil, fmt.Errorf("Error Reading state file: %w", err)
	}

	if checkExistingTodo(todos, title) {
		return nil, fmt.Errorf("Todo already exists: %s", title)
	}

	newTodo := Todo{
		Id:        len(todos) + 1,
		Title:     title,
		CreatedAt: time.Now(),
	}

	todos = append(todos, newTodo)

	if err := pkgs.WriteStateFile(todos); err != nil {
		return nil, fmt.Errorf("Error writing state file: %w", err)
	}

	return todos, nil
}

func ToggleDone(id int) (Todos, error) {
	todos, err := Read()

	if err != nil {
		return nil, fmt.Errorf("Error Reading state file: %w", err)
	}

	for i, t := range todos {
		if t.Id == id {
			todos[i].Completed = !todos[i].Completed
			todos[i].CreatedAt = time.Now()
		}
	}

	if err := pkgs.WriteStateFile(todos); err != nil {
		return nil, fmt.Errorf("Error writing state file: %w", err)
	}

	return todos, nil
}

func Delete(id int) (Todos, error) {
	todos, err := Read()
	if err != nil {
		return nil, fmt.Errorf("Error Reading state file: %w", err)
	}

	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
		}
	}

	if err := pkgs.WriteStateFile(todos); err != nil {
		return nil, fmt.Errorf("Error writing state file: %w", err)
	}

	return todos, nil
}
