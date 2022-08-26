package git

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// git remote -v
func TestFormatRemote(t *testing.T) {

	remotes := `origin  git@github.com:lgylgy/gitw.git (fetch)
origin  git@github.com:lgylgy/gitw.git (push)
upstream        git@github.com:upstream/gitw.git (fetch)
upstream        git@github.com:upstream/gitw.git (push)`

	expected := []string{
		"origin git@github.com:lgylgy/gitw.git",
		"upstream git@github.com:upstream/gitw.git",
	}

	formatted := formatRemotes(remotes)
	require.Equal(t, formatted, expected)
}

// git brancgh -v
func TestFormatCurrentBranch(t *testing.T) {

	remotes := `branch1  621c170a49 [ahead 49] Merge pull request #1234
* currentbranch   1a0ac73ba6 [ahead 25, behind 1] message commit.
branch3 d7c95e1449 [behind 123] message commit.`

	expected := `* currentbranch 1a0ac73ba6 [ahead 25, behind 1] message commit.`

	formatted := formatCurrentBranch(remotes)
	require.Equal(t, formatted, expected)
}
