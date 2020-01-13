package cmd

import (

	"os"

	"../db"
	"../initialiser"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your project directory for fast cd.",
	Run: func(cmd *cobra.Command, args []string) {
		tags,err := db.AllDir(initialiser.TagMasterData)
		if !CheckTagExists(groupTags,tags) && groupTags != initialiser.TagMasterData && groupTags != initialiser.TaskBucket  {
			d.Println("CRITICAL","Tag missing")
			os.Exit(1)
		}
		
		dirs, err := db.AllDir(groupTags)
		if err != nil {
			d.Println("CRITICAL","Something went wrong:", err)
			os.Exit(1)
		}
		if len(dirs) == 0 {
			d.Println("WARNING","You have no project directories saved! 🏖")
			return
		}
		d.Println("DEBUG","You have the project directories saved!")
		for i, dir := range dirs {
			d.Printf("%d. %s","DEBUG", i+1, dir.Value)
		}
	},
}

func init() {
	listCmd.Flags().StringVarP(&groupTags, "tag", "t", taskBucket, "Add a tag for your project directory")
	RootCmd.AddCommand(listCmd)
}

func CheckTagExists(a string, list []db.Dirs) bool {
    for _, b := range list {
        if b.Value == a {
            return true
        }
    }
    return false
}