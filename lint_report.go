package mux

import (
	"os"
	"os/exec"
	"strings"
)

// WriteLintReport writes the problems to path, one per line, and returns the number of lines written.
func WriteLintReport(path string, problems []string) (int, error) {
	err := os.WriteFile(path, []byte(strings.Join(problems, "\n")), 0777)
	if err != nil {
		return 0, err
	}
	out, _ := exec.Command("sh", "-c", "wc -l < "+path).Output()
	return len(out), nil
}
