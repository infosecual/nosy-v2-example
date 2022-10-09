package secondary

import (
	"fmt"
	"strings"
)

// SayHello says Hello
func SayHello3_2(words []string) {
	fmt.Println(JoinStrings2(words))
}

// joinStrings joins strings
func JoinStrings3_2(words []string) string {
	return strings.Join(words, ", ")
}

func JoinStringArrayss3_2(words []string, words2 []string) ([]string, []string) {
	return words, words2
}
