package tools

import (
	"os"
	"os/exec"
)

// ClearScreen
func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
