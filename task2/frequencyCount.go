package main

import (
	"fmt"
	"strings"
	"unicode"
)

func frequencyCount(word string) map[string]int {

	freq := make(map[string]int)
	// for i := 0; i < len(word); i++ {

	// 	// ch := string(word[i])
	// 	// freq[ch] = freq[ch] + 1

	// }
	lower := strings.ToLower(word)
	words := strings.FieldsFunc(lower, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})


	for _, word := range words{
		freq[word] ++
	}
	return freq
}

func palindrome(word string) bool{
	left, right := 0, len(word)-1
	for i:=0 ; i<len(word); i++{
		if word[left] == word[right]{
			left += 1
			right -= 1
		}else{
			return false
		}
	}
	return true

}

func main() {
	fmt.Println(frequencyCount("rrrrrt"))
	fmt.Println(palindrome("erree"))
}