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
    fmt.Println(JoinStrings(words))
}

// JoinStrings joins strings
func JoinStrings5(words []string) string {
    return strings.Join(words, ", ")
}

type StringToAdd struct {
        Name string
}

func (s StringToAdd) AppendFoo() string {
        return s.Name + "-foo!"
}

type IntToAdd struct {
        A int
}

func (a IntToAdd) AppendInt(b int) int {
        return a.A + b
}

func AddInts(int1, int2 int) int {
        return int1 + int2
}

func ReturnInts(int1, int2 int) (int, int) {
        return int1, int2
}

func AddIntsAndString(int1, int2 int, this string, that []string) (int, string, []string) {
        return int1 + int2, this, that
}

