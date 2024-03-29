package git

import (
	"bytes"
	"fmt"
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
	output, err := cmd(dir, args)
	if err != nil {
		return "", err
	}
	return formatCurrentBranch(output), err
}

func getRemotes(dir string) ([]string, error) {
	args := []string{"git", "remote", "--verbose"}
	output, err := cmd(dir, args)
	if err != nil {
		return nil, err
	}
	return formatRemotes(output), err
}

func getCommits(dir string) (string, error) {
	args := []string{"git", "log", "--graph", "--pretty=oneline", "--abbrev-commit", "--max-count=15"}
	output, err := cmd(dir, args)
	if err != nil {
		return "", err
	}
	return output, err
}

func getDiff(dir string, cached bool) (string, error) {
	args := []string{"git", "diff", "--name-only"}
	if cached {
		args = append(args, "--cached")
	}
	output, err := cmd(dir, args)
	if err != nil {
		return "", err
	}
	return output, err
}

func fetch(dir, remote string) (string, error) {
	args := []string{"git", "fetch", remote}
	output, err := cmd(dir, args)
	if err != nil {
		return "", err
	}
	return output, err
}

func resetHard(dir, remote, branch string) (string, error) {
	args := []string{"git", "reset", "--hard", fmt.Sprintf("%s/%s", remote, branch)}
	output, err := cmd(dir, args)
	if err != nil {
		return "", err
	}
	return output, err
}

func push(dir, remote, branch string) (string, error) {
	args := []string{"git", "push", remote, branch}
	output, err := cmd(dir, args)
	if err != nil {
		return "", err
	}
	return output, err
}
