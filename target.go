package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

type StringToAdd struct {
	Name string
}

type ComplexStruct struct {
	Name              string
	HexRepresentation []byte
}

func main() {
	hello := "Hello"
	world := "World"
	words := []string{hello, world}
	SayHello(words)
	s := StringToAdd{Name: "add to me "}
	fmt.Println(s.AppendFoo())
	Int := IntToAdd{A: 18446744073709551615}
	fmt.Println(Int.AppendInt(10))
	//c := ComplexStruct{}
	//c.SetNameViaHex("48656c6c6f20476f7068657221")
}

// SayHello says Hello
func SayHello(words []string) {
	fmt.Println(JoinStrings(words))
}

// JoinStrings joins strings
func JoinStrings(words []string) string {
	return strings.Join(words, ", ")
}

func AddInts(int1 IntToAdd, int2 uint64) uint64 {
	return int1.A + int2
}

func (s StringToAdd) AppendFoo() string {
	return s.Name + "-foo!"
}

type IntToAdd struct {
	A uint64
}

func (a IntToAdd) AppendInt(b uint64) uint64 {
	return a.A + b
}

func (c ComplexStruct) SetNameViaHex(src []byte) {

	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		panic(err)
	}
	c.Name = string(dst[:n])
	fmt.Printf("ComplexStruct name set to %s", c.Name)
}
