package main

import (
	"fmt"
)

func main() {
	var val int64
	var mask int64
	var i int
	var op int
	// 1-установка бита
	// 0-сброс бита
	fmt.Printf("Введите желаемую операцию")
	fmt.Scanln(&op)
	fmt.Printf("Введите номер бита")
	fmt.Scanln(&i)
	mask = 1
	// Формируется маска, сдвигается влево число 1 (1 << 3 = 1000)
	mask = mask << i
	val = 23456665467655
	fmt.Printf("value :%b\nmask :%b\n", val, mask)
	if op == 1 {
		// При установке бита используется операция побитового "или", при условии, что бит уже установлен,
		// иначе значение val не изменится
		val = val | mask
	} else {
		// При сбросе бита маска сначала инвертируется, а затем применяется операция побитового "и", при условии, что бит установлен,
		// иначе значение val не изменится
		val = val & ^mask
	}
	fmt.Printf("new val :%b\n", val)
}
