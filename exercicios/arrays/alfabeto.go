package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	textAsNumber := encryptText("Mamaco")
	fmt.Println(textAsNumber)
}

func encryptText(text string) string {
	alphabet := [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	text = strings.ToLower(text)
	splitedText := strings.Split(text, "")
	sequence := ""

	for _, letter := range splitedText {
		for j, character := range alphabet {
			index := j + 1
			indexAsString := strconv.Itoa(index)
			if letter == character {
				sequence += indexAsString
			}
		}
	}
	return sequence
}
