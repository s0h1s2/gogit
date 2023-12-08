package main

import (
	"gogit/internal"
	"log"

	"github.com/spf13/cobra"
)

// var args struct {
// 	Init       *internal.InitCmd       `arg:"subcommand:init"`
// 	HashObject *internal.HashObjectCmd `arg:"subcommand:hash-object"`
// 	HashCat    *internal.HashCatCmd    `arg:"subcommand:cat"`
// 	WriteTree  *internal.WriteTreeCmd  `arg:"subcommand:write-tree"`
// }

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
		internal.InitializeRepository()
	},
}

func main() {
	rootCmd.AddCommand(initCmd)
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err.Error())
	}
	// p := arg.MustParse(&args)
	// if p.Subcommand() == nil {
	// 	p.WriteHelp(os.Stderr)
	// 	return
	// }
	// if args.Init != nil {
	// 	internal.InitializeRepository()
	// 	return
	// }
	// switch {
	// case args.HashObject != nil:
	// 	{
	// 		internal.CreateBlobHashObject(args.HashObject)
	// 	}
	// case args.HashCat != nil:
	// 	{
	// 		print(internal.HashCat(args.HashCat, internal.ObjTag(args.HashCat.Type)))
	// 	}
	// case args.WriteTree != nil:
	// 	{
	// 		internal.WriteTree()
	// 	}
	// }
}
