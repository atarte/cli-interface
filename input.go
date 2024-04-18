package main

import (
	"errors"
	"fmt"
	"strconv"
)

// ConfirmInput
func ConfirmInput(question string) bool {
	fmt.Printf("%s (y/n)\n", question)
	// input =
	return true
}

type ChoiceElement struct {
	Text  string
	Input string
	Exec  func()
}

// isChoiceInputFromList
func isChoiceInputFromList(choiceList []ChoiceElement) bool {
	for _, choice := range choiceList {
		if choice.Input == "" {
			return false
		}
	}
	return true
}

// ChoiceInput
func ChoiceInput(question string, choiceList []ChoiceElement) error {
	fmt.Printf("%s\n", question)

	if !isChoiceInputFromList(choiceList) {
		return errors.New("Missing `input`field in choice element")
	}

	for _, choice := range choiceList {
		fmt.Printf("(%s) - %s\n", choice.Input, choice.Text)
	}

	var input string
	fmt.Scanln(&input)

	for _, choice := range choiceList {
		if input == choice.Input {
			choice.Exec()
			break
		}
	}

	return nil
}

// ChoiceNumberInput
func ChoiceNumberInput(question string, choiceList []ChoiceElement, offset bool) {
	fmt.Printf("%s\n", question)

	for i, choice := range choiceList {
		if offset {
			i++
		}

		fmt.Printf("(%d) - %s\n", i, choice.Text)
	}

	var input string
	fmt.Scanln(&input)

	for i, choice := range choiceList {
		if offset {
			i++
		}

		if input == strconv.Itoa(i) {
			choice.Exec()
			break
		}
	}
}
