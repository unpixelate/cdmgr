package cmd

import (
	"fmt"
	"strconv"
	"os"

	"../initialiser"
	"../db"

	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete a project directory in database.",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		dirlist := ReturnDirectoriesToBeDeleted(taskBucket,ids)
		allBuckets , err := db.AllDir(initialiser.TagMasterData)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		for _,bkt := range(allBuckets){
			DeleteFromDirectory(bkt.Value,dirlist)
		} 
		DeleteFromDirectory(initialiser.TaskBucket, dirlist)
		
	},
}

func init() {
	delCmd.Flags().StringVarP(&groupTags, "tag", "t", taskBucket, "Add a tag for your project directory")
	RootCmd.AddCommand(delCmd)
}


func ReturnDirectoriesToBeDeleted(taskBucket string,ids []int)[]string{
	tasks, err := db.AllDir(taskBucket)
	if err != nil {
		fmt.Println("Something went wrong:", err)
		os.Exit(2)
	}
	var dirList []string
	for _, id := range ids {
		if id <= 0 || id > len(tasks) {
			d.Println("DEBUG","Invalid task number:", id)
			continue
		}
		task := tasks[id-1]
		dirList = append(dirList,task.Value)
	}
	return dirList
}

func DeleteFromDirectory(taskBucket string,dirlist []string){
	dirs, err := db.AllDir(taskBucket)
	if err != nil {
		fmt.Println("Something went wrong:", err)
		os.Exit(2)
	}
	for key, dir := range dirs {
		for _,targetDirectory := range(dirlist){
			if dir.Value == targetDirectory{
				err := db.DeleteDir(dir.Key, taskBucket)
				if err != nil {
					d.Printf("Failed to mark \"%d\" as completed. Error: %s","CRITICAL", key, err)
				} else {
					d.Printf("Deleted %s from \"%s\".","CRITICAL", dir.Value,taskBucket)
				}
				break
			}
		}
	}
}

