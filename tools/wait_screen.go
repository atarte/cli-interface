package tools

import "fmt"

// WaitScreen
func WaitScreen() {
	fmt.Println("")
	fmt.Println("(Press enter to return...)")
	fmt.Scanln()
}
