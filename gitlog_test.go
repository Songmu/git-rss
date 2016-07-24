package gitrss

import (
	"os"
	"testing"

	"github.com/k0kubun/pp"
)

func TestParseGitlog(t *testing.T) {
	fp, err := os.Open("testdata/gitlog.txt")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer fp.Close()
	r := parseGitlog(fp)
	pp.Print(r)
}
