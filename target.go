package main

import (
	"fmt"
	"strings"
)

type StringToAdd struct {
        Name string
}

func main() {
    hello := "Hello"
    world := "World"
    words := []string{hello, world}
    SayHello(words)
    s := StringToAdd{Name:"add to me "}
    fmt.Println(s.AppendFoo())
    Int := IntToAdd{A:5}
    fmt.Println(Int.AppendInt(10))
}

// SayHello says Hello
func SayHello(words []string) {
    fmt.Println(JoinStrings(words))
}

// JoinStrings joins strings
func JoinStrings(words []string) string {
    return strings.Join(words, ", ")
}

func AddInts(int1, int2 int) int {
        return int1 + int2
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
