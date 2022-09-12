package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "qweErty"
	str2 := "qqwerty"
	str1Utf := "абБвгд"
	str2Utf := "аабвгд"
	fmt.Println("str1: ", isUniqueChars2(str1))
	fmt.Println("str2: ", isUniqueChars2(str2))
	fmt.Println("str1Utf: ", isUniqueChars2(str1Utf))
	fmt.Println("str2Utf: ", isUniqueChars2(str2Utf))
}

func isUniqueChars2(str string) bool {
	strings.ToLower(str)
	m := make(map[rune]bool)
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		val := runes[i]
		if m[val] { //символ уже был найден в строке
			return false
		}
		m[val] = true
	}
	return true
}
