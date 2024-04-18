package main

import (
	"fmt"
	"os"
	"os/exec"
)

// ClearScreen
func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// WaitScreen
func WaitScreen() {
	fmt.Println("")
	fmt.Println("(Press enter to return...)")
	fmt.Scanln()
}
