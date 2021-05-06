package cmd

import (
	"fmt"
	"strconv"
	"todo/db"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks tasks as completed",

	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println(err)
			} else {
				ids = append(ids, id)
			}
		}
		//fmt.Println(ids)
		tasks, err := db.AllTask()
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid Number")
				continue
			}

			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Println(err)

			} else {
				fmt.Println("Succesfully marked as completed")
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)

}
