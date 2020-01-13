package main

import (
	"fmt"
	"os"

	"./src/cmd"
	"./src/db"
	"./src/initialiser"
)

func main() {
	initialiser.InitiiseRootFolder()
	dbPath := initialiser.ReturnDbPath()
	must(db.Init(dbPath))
	must(db.CreateBucket(initialiser.TaskBucket))
	must(db.CreateBucket(initialiser.TagMasterData))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
