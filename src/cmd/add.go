package cmd

import (
	"log"
	"os"

	"../db"
	"../initialiser"
	"../tags"

	"github.com/spf13/cobra"
)

var groupTags string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a project directory to your directory list.",
	Run: func(cmd *cobra.Command, args []string) {
		// task := strings.Join(args, " ")
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		d.Println("DEBUG","Current directory is : "+dir)
		// handles the main event
		createDir(dir, taskBucket)
		// handles additional tag

		if groupTags != "" && groupTags != taskBucket {
			createDir(dir, groupTags)
			if checkExisitngTags(groupTags) == false && groupTags != initialiser.TaskBucket {
				db.CreateDirectory(groupTags, initialiser.TagMasterData)
			}
		}

	},
}

func init() {
	addCmd.Flags().StringVarP(&groupTags, "tag", "t", taskBucket, "Add a tag for your project directory")
	RootCmd.AddCommand(addCmd)
}

// function handles both main and tags
func createDir(dir string, groupTags string) {
	err := tags.Init(groupTags)
	if err != nil {
		d.Println("CRITICAL", err.Error())
		return
	}
	_, err = db.CreateDirectory(dir, groupTags)
	d.Printf("Added \"%s\" to your %s db.", "DEBUG", dir, groupTags)
	if err != nil {
		d.Printf("Something went wrong in %s database: %s", "CRITICAL", groupTags, err)
		return
	}
}

func checkExisitngTags(groupTags string) bool {
	tags, err := db.AllDir(initialiser.TagMasterData)
	if err != nil {
		d.Println("WARNING", "Error")
	}
	for _, v := range tags {
		if groupTags == v.Value {
			return true
		}
	}
	return false
}
