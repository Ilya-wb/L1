package main

import "fmt"

func main() {
	str := "главрыба"
	fmt.Println(Reverse(str))
	str = "qwerty"
	fmt.Println(Reverse(str))
}

func Reverse(s string) string {
	// Использование рун позволяет считывать строку по отдельному символу, а не байту
	// Это позволяет использовать как asci(1б), так и unicode(2б)
	runes := []rune(s)
	// Выполение реверса массива
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		// Элементы последовательно взаимозаменяют друг-друга
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
