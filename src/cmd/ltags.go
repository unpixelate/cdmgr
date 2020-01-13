package cmd

import (
	"os"

	"../db"
	"../initialiser"
	"github.com/spf13/cobra"
)

var ltags = &cobra.Command{
	Use:   "tags",
	Short: "Lists all tags.",
	Run: func(cmd *cobra.Command, args []string) {
		tags, err := db.AllDir(initialiser.TagMasterData)
		if err != nil {
			d.Println("WARNING", "Something went wrong:", err)
			os.Exit(1)
		}
		if len(tags) == 0 {
			d.Println("WARNING", "You have no tags. Consider adding one for faster lookup!")
			return
		}
		d.Println("DEBUG", "Here are your tags!")
		for i, dir := range tags {
			d.Printf("%d. %s", "DEBUG", i+1, dir.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(ltags)
}
