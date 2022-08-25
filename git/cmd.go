package git

import (
	"bytes"
	"os/exec"
)

func cmd(dir string, args []string) (string, error) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = dir

	out, err := cmd.CombinedOutput()
	if err == nil && len(out) > 0 {
		return string(bytes.TrimSpace(out)), nil
	}
	return "", err
}

func getCurrentBranch(dir string) (string, error) {
	args := []string{"git", "branch", "--verbose"}
	return cmd(dir, args)
}

func getRemotes(dir string) (string, error) {
	args := []string{"git", "remote", "--verbose"}
	return cmd(dir, args)
}
