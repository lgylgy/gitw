package git

import (
	"bufio"
	"fmt"
	"strings"
)

func formatRemotes(text string) []string {
	scanner := bufio.NewScanner(strings.NewReader(text))
	result := []string{}

	// map to remove duplicate remote: fetch / push
	remotes := map[string]bool{}
	for scanner.Scan() {
		// remove duplicate whitespaces
		text := strings.Join(strings.Fields(strings.TrimSpace(scanner.Text())), " ")
		parts := strings.Split(text, " ")
		if len(parts) >= 3 {
			label := fmt.Sprintf("%s %s", parts[0], parts[1])
			_, ok := remotes[label]
			if !ok {
				result = append(result, label)
			}
			remotes[label] = true
		}
	}
	return result
}

func formatCurrentBranch(text string) string {
	scanner := bufio.NewScanner(strings.NewReader(text))
	for scanner.Scan() {
		// remove duplicate whitespaces
		text := strings.Join(strings.Fields(strings.TrimSpace(scanner.Text())), " ")
		parts := strings.Split(text, " ")
		if len(parts) > 0 && parts[0] == "*" {
			return text
		}
	}
	return ""
}
