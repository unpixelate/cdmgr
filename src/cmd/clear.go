package cmd

import (
	"fmt"
	"strconv"


	"../initialiser"
	"../db"

	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Delete a optional database.",
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
		tags, err := db.AllDir(initialiser.TagMasterData)

		for _ , id := range(ids){
			
			err := db.DeleteDir(id, initialiser.TagMasterData)
			if err != nil {
				d.Printf("Failed to delete bucket \"%d\". Error: %s","CRITICAL", err,id)
			} else {
				d.Printf("Deleted bucket %d","WARNING",id, dirlist[id])
			}
			err = db.DeleteBucket(dirlist[id])
			if err != nil{
				d.Println("Something went wrong... unable to delete bucket")
				return
			}
		} 
	},
}

func init() {
	//clearCmd.Flags().StringVarP(&groupTags, "tag", "t", taskBucket, "Delete tag bucket")
	//clearCmd.MarkFlagRequired("tag")
	RootCmd.AddCommand(clearCmd)
}
