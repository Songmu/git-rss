package gitrss

import "os/exec"

type git struct {
	gitPath string
}

type gitrepo struct {
	repoPath string
	*git
}

func newGitrepo(path string) *gitrepo {
	return &gitrepo{
		repoPath: path,
		git:      &git{},
	}
}

func (g *git) prog() string {
	if g.gitPath != "" {
		return g.gitPath
	}
	return "git"
}

func (gr *gitrepo) getRepoPath() string {
	if gr.repoPath == "" {
		gr.repoPath = "."
	}
	return gr.repoPath
}

func (gr *gitrepo) cmd(args ...string) *exec.Cmd {
	argv := []string{"-C", gr.getRepoPath()}
	argv = append(argv, args...)
	return exec.Command(gr.prog(), argv...)
}
