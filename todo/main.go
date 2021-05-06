package main

import (
	"fmt"
	"os"
	"path/filepath"
	"todo/cmd"
	"todo/db"

	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))

	cmd.RootCmd.Execute()

}

func must(err error) {

	if err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}
}
