package secondary

import "fmt"
import "strings"

func main() {
    hello := "Hello"
    world := "World"
    words := []string{hello, world}
    SayHello(words)
}

// SayHello says Hello
func SayHello2(words []string) {
    fmt.Println(JoinStrings(words))
}

// JoinStrings joins strings
func JoinStrings2(words []string) string {
    return strings.Join(words, ", ")
}
