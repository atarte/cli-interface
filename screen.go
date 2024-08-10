package cliinterface

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// ClearScreen
func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// WaitScreen
func WaitScreen() {
	fmt.Println("")
	fmt.Println("(Press enter to return...)")
	fmt.Scanln()
}
