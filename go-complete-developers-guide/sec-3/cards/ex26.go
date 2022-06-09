package main

import (
	"fmt"
	"strings"
)

func main() {

	print_first_and_last("This is cool stuff and all")
}

func break_words(sentence string) []string {
	words := strings.Split(sentence, " ")
	return words
}

func print_first_word(words []string) {
	x := words[0]
	fmt.Println(x)
}

func print_last_word(words []string) {
	word := words[len(words)-1]
	fmt.Println(word)
}

func print_first_and_last(sentence string) {
	words := break_words(sentence)
	print_first_word(words)
	print_last_word(words)
}
