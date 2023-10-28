package main

import (
	"fmt"

	"github.com/atarte/console-screen/tools"
	"github.com/atarte/console-screen/utils"
)

func main() {
	// fmt.Printf("\atest\n")
	// fmt.Printf("%0113d\n", 1)

	size := utils.GetWindSize()
	fmt.Println("Windiws size :", size)

	count := 11000
	for i := 0; i < count; i++ {
		tools.LoadingBar(i+1, count)

		// Wait a bit
		// time.Sleep(1 * time.Second)
	}
}
