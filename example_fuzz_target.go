package example

import (
	"fmt"
	"os"
	"bufio"
)

/*
# built_in_types
'string', '[]byte', 'int', 'int8', 'int16', 'int32', 'rune',
'int64', 'uint', 'uint8', 'byte', 'uint16', 'uint32',
'uint64', 'float32', 'float64', 'bool']
*/

type complex_struct struct {
	a string
	b []byte
	c int
	d int8
	e int16
	f int32
	g rune
	h int64
	i uint
	j uint8
	k byte
	l uint16
	m uint32
	n uint64
	o float32
	p float64
	q bool
}

func guess_int(x int) bool {
	if x == 5 {
		return true
	}
	return false
}

func guess_string(x string) bool {
        if x == "this is a long string, too hard to bruteforce quickly\n" {
                return true
        }
        return false
}


func main() {

	// test int
	var x int
	fmt.Println("enter a number:")
	if _, err := fmt.Scan(&x); err != nil {
		fmt.Println("  Scan for i failed, due to ", err)
		return
	}

	if guess_int(x ) == false {
		fmt.Println("failed int check")
		return
	}
	fmt.Println("success!")

        // test string
	var y string
        fmt.Println("enter a string:")
	reader := bufio.NewReader(os.Stdin)
	y, _ = reader.ReadString('\n')
        if guess_string(y) == false {
                fmt.Println("failed string check")
                return
        }
        fmt.Println("success!")



}


