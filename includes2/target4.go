package quadrary

import "fmt"
import "strings"

func main() {
    hello := "Hello"
    world := "World"
    words := []string{hello, world}
    SayHello(words)
}

// SayHello says Hello
func SayHello4(words []string) {
    fmt.Println(JoinStrings(words))
}

// JoinStrings joins strings
func JoinStrings4(words []string) string {
    return strings.Join(words, ", ")
}
