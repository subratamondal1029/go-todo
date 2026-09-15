package todo

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/subratamondal1029/goTodo/internal/database"
)

func PrintDataTable(todos *[]database.Todo) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer w.Flush()

	fmt.Fprintln(w, "ID\tTITLE\tCOMPLETED")
	for _, todo := range *todos {
		fmt.Fprintf(w, "%d\t%s\t%t\n",
			todo.ID,
			todo.Title,
			todo.Completed,
		)
	}
}
