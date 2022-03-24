package includes

/*
# built_in_types
'string', '[]byte', 'int', 'int8', 'int16', 'int32', 'rune',
'int64', 'uint', 'uint8', 'byte', 'uint16', 'uint32',
'uint64', 'float32', 'float64', 'bool']
*/

func Guess_int(x int) bool {
	if x == 5 {
		panic("CRASH")
	}
	return false
}

func Guess_hard_string(x string) bool {
        if x == "this is a long string, too hard to bruteforce quickly\n" {
                panic("CRASH")
        }
        return false
}

func Guess_easy_string(x string) bool {
        if x == "a" {
                panic("CRASH")
        }
        return false
}

func Guess_int32(x int32) bool {
        var y int32
	y = 500
	if x == y {
                panic("CRASH")
        }
        return false
}


