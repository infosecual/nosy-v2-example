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
func SayHello3_2(words []string) {
    fmt.Println(joinStrings(words))
}

// joinStrings joins strings
func joinStrings3_2(words []string) string {
    return strings.Join(words, ", ")
}

func joinStringArrayss3_2(words []string, words2 []string) (string, string) {
    return words, words2
}

