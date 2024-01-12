package main

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Gogit",
	Short: "Gogit is git like vsc",
	Long:  `Gogit is a git like vsc written from scratch for educational purpose`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize repoistory",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func main() {
	rootCmd.AddCommand(initCmd)
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err.Error())
	}
}
