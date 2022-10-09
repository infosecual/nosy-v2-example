package quadrary

import (
	"fmt"
	"strings"
)

func main() {
	hello := "Hello"
	world := "World"
	words := []string{hello, world}
	SayHello4(words)
}

// SayHello says Hello
func SayHello4(words []string) {
	fmt.Println(JoinStrings4(words))
}

// JoinStrings joins strings
func JoinStrings4(words []string) string {
	return strings.Join(words, ", ")
}
