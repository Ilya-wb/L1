package main

import "fmt"

// Конвертируем в байты и руны и измеряем длину соответственно
// Если количество байт не равно количеству символов, значит используется юникод, а в нем каждый символ
// имеет размер 2 байта
// Срез v[:n] берет именно количество байт
// В данной реализации строка может быть как на кириллице, так и на латинице
func slice(str string, n int) {
	if len([]byte(str)) != len([]rune(str)) {
		n *= 2
	}
	justString := str[:n]
	fmt.Println(justString)
}

func main() {
	v := "eaztrdhzdrhdzthdtfhfdtxhtfxhxtfhfhfth"
	a := "аволимтраолитваиолатиаолитасиоастиоал"
	slice(v, 10)
	slice(a, 10)
}
