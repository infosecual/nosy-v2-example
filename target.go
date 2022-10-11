package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

type StringToAdd struct {
	Name string
}

type ComplexStruct struct {
	Name              string
	HexRepresentation string
	RandomByteData    []byte
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
	c := ComplexStruct{Name: "test struct ",
		HexRepresentation: "0x48656c6c6f20476f7068657221",
		RandomByteData:    []byte{2, 3, 5, 7, 11, 13}}
	c.RepeatNameXTimes(2)
	c.DecodeHex()
	c.Print5thByte()
	c.DivideXByteByY(3, 2)
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

func (c ComplexStruct) RepeatNameXTimes(x int) {

	output := bytes.Buffer{}
	output.WriteString(strings.Repeat(c.Name, x*2))
	fmt.Println("Repeating struct name", x-1, "times")
	fmt.Println(output.String())
}

// Decode decodes a hex string with 0x prefix.
func (c ComplexStruct) DecodeHex() {
	dec, err := hexutil.Decode(c.HexRepresentation)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decoded:", dec)
}

func (c ComplexStruct) Print5thByte() {
	fmt.Println(c.RandomByteData[5])
}

func (c ComplexStruct) DivideXByteByY(x int, y int) {
	fmt.Println(int(c.RandomByteData[x]) / y)
}
