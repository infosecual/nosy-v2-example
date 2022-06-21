package tertiary

import "fmt"
import "strings"

func main() {
    hello := "Hello"
    world := "World"
    words := []string{hello, world}
    SayHello(words)
}

// SayHello says Hello
func SayHello5(words []string) {
    fmt.Println(joinStrings(words))
}

// joinStrings joins strings
func joinStrings5(words []string) string {
    return strings.Join(words, ", ")
}

type StringToAdd struct {
        Name string
}

func (s StringToAdd) appendFoo() string {
        return s.Name + "-foo!"
}

type IntToAdd struct {
        A int
}

func (a IntToAdd) appendInt(b int) int {
        return a.A + b
}

func addInts(int1, int2 int) int {
        return int1 + int2
}

func returnInts(int1, int2 int) (int, int) {
        return int1, int2
}

func addIntsAndString(int1, int2 int, this string, that []string) (int, string, []string) {
        return int1 + int2, this, that
}

