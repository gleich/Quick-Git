package git

import (
	"io/ioutil"
	"os"

	"github.com/Matt-Gleich/statuser/v2"
)

// ValidateCWD ... Make sure that the current working directory is a git repository
func ValidateCWD() {
	contents, err := ioutil.ReadDir(".")
	if err != nil {
		statuser.Error("Failed to list current directory", err, 1)
	}

	var found bool
	for _, fileOrDir := range contents {
		if fileOrDir.Name() == ".git" && fileOrDir.IsDir() {
			found = true
		}
	}

	if !found {
		wd, err := os.Getwd()
		if err != nil {
			statuser.Error("Failed to get current directory path", err, 1)
		}
		statuser.ErrorMsg("The current directory ("+wd+") is not a git repository", 1)
	}
}
