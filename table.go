package cliinterface

import (
	"fmt"
)

type TableElement struct {
	Header []string
	Body   [][]string
}

// getLongestStringFromCollum
func (t *TableElement) getLongestStringFromCollum(index int) int {
	max_len := len(t.Header[index])

	for _, col := range t.Body {
		if len(col[index]) > max_len {
			max_len = len(col[index])
		}
	}

	return max_len
}

// justifyText
func justifyText(text string, max_len int) {
	justify := max_len - len(text) + 3
	for i := 0; i < justify; i++ {
		fmt.Print(" ")
	}
}

// Table
func Table(table TableElement) {
	// Header display
	for i, header := range table.Header {
		tab_len := table.getLongestStringFromCollum(i)
		fmt.Printf("%s", header)
		justifyText(header, tab_len)
	}
	fmt.Println("")

	// Body display
	for _, col := range table.Body {
		for i, row := range col {
			tab_len := table.getLongestStringFromCollum(i)
			fmt.Printf("%s", row)
			justifyText(row, tab_len)
		}
		fmt.Println("")
	}
}
