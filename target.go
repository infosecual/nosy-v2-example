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
    fmt.Println(s.appendFoo())
    Int := IntToAdd{A:5}
    fmt.Println(Int.appendInt(10))
}

// SayHello says Hello
func SayHello(words []string) {
    fmt.Println(joinStrings(words))
}

// joinStrings joins strings
func joinStrings(words []string) string {
    return strings.Join(words, ", ")
}

func addInts(int1, int2 int) int {
        return int1 + int2
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
