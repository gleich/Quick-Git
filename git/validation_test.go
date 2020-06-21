package git

import (
	"os"
	"testing"

	"github.com/Matt-Gleich/statuser/v2"
)

func TestAValidateCWD(t *testing.T) {
	_, err := os.Stat(".git")
	if os.IsNotExist(err) {
		err := os.MkdirAll(".git", 0755)
		if err != nil {
			statuser.Error("Failed to make directory: .git", err, 1)
		}
	}
	ValidateCWD()
}
