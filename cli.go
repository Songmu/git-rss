package gitrss

import (
	"io"
	"log"
)

// CLI is struct for command line tool
type CLI struct {
	OutStream, ErrStream io.Writer
}

// Run the ghch
func (cli *CLI) Run(argv []string) int {
	log.SetOutput(cli.ErrStream)
	run()
	return 0
}
