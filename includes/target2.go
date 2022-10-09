package secondary

import (
	"fmt"
	"strings"
)

func main() {
	hello := "Hello"
	world := "World"
	words := []string{hello, world}
	SayHello2(words)
}

// SayHello says Hello
func SayHello2(words []string) {
	fmt.Println(JoinStrings2(words))
}

// JoinStrings joins strings
func JoinStrings2(words []string) string {
	return strings.Join(words, ", ")
}
