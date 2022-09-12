package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	// Разбиение строки str по разделителю "пробел" на подстроки
	// Срез words будет содержать отдельные подстроки(слова)
	words := strings.Split(str, " ")
	// Реверс массива words
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	fmt.Println(words)
}
