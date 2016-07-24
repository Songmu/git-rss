package main

import (
	"os"

	"github.com/Songmu/git-rss"
)

func main() {
	os.Exit((&gitrss.CLI{ErrStream: os.Stderr, OutStream: os.Stdout}).Run(os.Args[1:]))
}
