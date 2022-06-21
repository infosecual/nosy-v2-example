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
    fmt.Println(joinStrings(words))
}

// joinStrings joins strings
func joinStrings2(words []string) string {
    return strings.Join(words, ", ")
}
