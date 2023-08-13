package main

import (
	"fmt"
	"unicode/utf8"
)

var p = fmt.Println

func main() {
	var text = "HelloWorld\n  HellooooooWorlddd  \nHellooo"
	var words int = 0
	var lines int = 0

	//count characters
	p("Number of characters, applying a library: ", utf8.RuneCountInString(text))
	p("Number of characters: ", len(text))

	//numbers of words and numbers of lines
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			words++
		}
		if text[i] == '\n' {
			lines++
			i++
		}
	}
	p("The number of words corresponds to: ", words, "\nThe number of lines corresponds to: ", lines)

}
