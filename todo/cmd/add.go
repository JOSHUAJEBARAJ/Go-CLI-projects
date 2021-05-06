package cmd

import (
	"fmt"
	"strings"
	"todo/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add command to the task",
	Run: func(cmd *cobra.Command, args []string) {
		// for i, arg := range args {
		// 	fmt.Println(i, arg)
		// }
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
