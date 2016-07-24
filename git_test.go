package gitrss

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/Songmu/gitmock"
)

func TestGitrepo(t *testing.T) {
	gm, err := gitmock.New("")
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer os.RemoveAll(gm.RepoPath())
	gr := newGitrepo(gm.RepoPath())
	cmd := gr.cmd("version")
	var b bytes.Buffer
	cmd.Stdout = &b
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if !strings.HasPrefix(b.String(), "git version") {
		t.Errorf("git version failed: %s", b.String())
	}
}
