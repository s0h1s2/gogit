package main

import (
	"gogit/internal"
	"os"

	"github.com/alexflint/go-arg"
)

var args struct {
	Init       *internal.InitCmd       `arg:"subcommand:init"`
	HashObject *internal.HashObjectCmd `arg:"subcommand:hash-object"`
	HashCat    *internal.HashCatCmd    `arg:"subcommand:cat"`
	WriteTree  *internal.WriteTreeCmd  `arg:"subcommand:write-tree"`
}

func main() {
	p := arg.MustParse(&args)
	if p.Subcommand() == nil {
		p.WriteHelp(os.Stderr)
		return
	}
	if args.Init != nil {
		internal.InitializeRepository()
		return
	}
	switch {
	case args.HashObject != nil:
		{
			internal.CreateBlobHashObject(args.HashObject)
		}
	case args.HashCat != nil:
		{
			internal.HashCatBlob(args.HashCat)
		}
	case args.WriteTree != nil:
		{
			internal.WriteTree()
		}
	}
}
