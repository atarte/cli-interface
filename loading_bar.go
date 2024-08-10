package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/atarte/cli-interface/utils"
)

// LoadingBar
func LoadingBar(num int, max int) {
	// Get the bar size by using the GetWindSize function
	bar_size := utils.GetWindSize().Col - 6 - 2*len(strconv.FormatInt(int64(max), 10))

	//
	current_bar := int(math.Ceil(float64((num * bar_size) / max)))

	//
	bar := ""
	for i := 0; i < current_bar-1; i++ {
		bar += "="
	}

	//
	if num != 0 {
		if num == max {
			bar += "="
		} else {
			bar += ">"
		}
	}

	//
	for i := 0; i < bar_size-current_bar; i++ {
		bar += " "
	}

	//
	fmt.Printf("\r[%s] (%d/%d)", bar, num, max)
}
