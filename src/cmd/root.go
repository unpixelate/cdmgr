package cmd

import (
	"../display"
	"../initialiser"
	"github.com/spf13/cobra"
)

const taskBucket = initialiser.TaskBucket

var Userstatus string

// use d as a global variable to store user related settings ;D
var d display.Display

var RootCmd = &cobra.Command{
	Use:   "cdmanager",
	Short: "Cdmanager is a directory bookmark manager",
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&Userstatus, "log", "DEBUG", "Level of logs")
}

func initConfig() {
	statusCode, _ := RootCmd.Flags().GetString("log")
	if statusCode != "" {
		(&d).Globalstatus = statusCode
	}
}
