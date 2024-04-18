package main

import (
	"fmt"

	"github.com/atarte/console-screen/tools"
)

func main() {
	fmt.Printf("\atest\n")
	//
	//
	// fmt.Printf("%0113d\n", 1)

	// size := utils.GetWindSize()
	// fmt.Println("Windiws size :", size)

	// count := 11000
	// for i := 0; i < count; i++ {
	// 	tools.LoadingBar(i+1, count)

	// 	// Wait a bit
	// 	// time.Sleep(1 * time.Second)
	// }
	//

	// tools.ChoiceNumberInput("Es ce que la vie est cool", []tools.ChoiceElement{
	// 	{
	// 		Text:  "Text tres cool",
	// 		Input: "a",
	// 		Exec: func() {
	// 			fmt.Println("zzz")
	// 		},
	// 	},
	// 	{
	// 		Text:  "Test 2ème du nom",
	// 		Input: "b",
	// 	},
	// }, true)
	//

	tools.Table(tools.TableElement{
		Header: []string{
			"Header1", "Header2", "Header3",
		},
		Body: [][]string{
			{"a", "b", "c"},
			{"1", "2", "3"},
			{".", " ", "!"},
		},
	})

}
