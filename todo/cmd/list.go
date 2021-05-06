package cmd

import (
	"fmt"
	"os"
	"todo/db"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your task",

	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTask()
		if err != nil {
			fmt.Println("Something went wrong", err)
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("There is no task to complete")
			return
		}
		for i, task := range tasks {
			fmt.Printf("%d %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

}
