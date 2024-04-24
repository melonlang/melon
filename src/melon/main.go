package main

import (
	"melon/lexer"
	"os"
)

func main() {
	bytes, _ := os.ReadFile("./examples/main.mln")
	tokens := lexer.Tokenize(string(bytes))

	for _, token := range tokens {
		token.Debug()
	}
}
