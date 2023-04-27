package shadow_go

import (
	"os/exec"
	"strings"
)

func Branch() string {
	branch, _ := gitCommand{}.currentBranch()
	return branch
}

func Tag() string {
	tag, _ := gitCommand{}.currentTag()
	return tag
}

func GitClean() bool {
	clean, _ := gitCommand{}.clean()
	return clean
}

func GitStatusFile() string {
	branch, _ := gitCommand{}.statusFile()
	return branch
}

func cmdRun(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	var out strings.Builder
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}

type gitCommand struct{}

func (g gitCommand) currentBranch() (string, error) {
	out, err := cmdRun("git", []string{"symbolic-ref", "--short", "HEAD"}...)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(out), nil
}

func (g gitCommand) currentTag() (string, error) {
	out, err := cmdRun("git", []string{"tag", "-l", "--contains", "HEAD"}...)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(out), nil
}

func (g gitCommand) lastTag() (string, error) {
	out, err := cmdRun("git", []string{"describe", "--tags", "--abbrev=0", "HEAD"}...)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(out), nil
}

func (g gitCommand) clean() (bool, error) {
	out, err := cmdRun("git", []string{"status", "--porcelain"}...)
	if err != nil {
		return false, err
	}

	for _, r := range strings.Split(out, "\n") {
		if len(strings.TrimSpace(r)) != 0 {
			return false, nil
		}
	}
	return true, nil
}

func (g gitCommand) statusFile() (string, error) {
	dirty := gitStatusFiles([]string{"status", "--porcelain"}, "^\\sM.")
	stage := gitStatusFiles([]string{"status", "--porcelain", "--untracked-files=all"}, "^[A|M|D|R]")

	return filterGitDirtyStage(dirty, stage), nil
}

func filterGitDirtyStage(dirtyFiles, stagedFiles []string) string {
	ret := strings.Builder{}
	for _, file := range dirtyFiles {
		ret.WriteString("  * ")
		ret.WriteString(file)
		ret.WriteString(" (dirty)\n")
	}
	for _, file := range stagedFiles {
		ret.WriteString("  * ")
		ret.WriteString(file)
		ret.WriteString(" (staged)\n")
	}
	return ret.String()
}
