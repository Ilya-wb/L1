package main

import (
	"fmt"
	"time"
)

var sum int

// Возведение числа в квадрат
func square(num int) {
	//time.Sleep(100)
	//val := num * num
	//append(sumItems, val)
	sum += num * num
	fmt.Println("sum", sum)
	fmt.Println(num * num)
}

// В данной функции n раз запускается горутина, где n-количество элементов среза nums
func do() {
	nums := [...]int{2, 4, 6, 8, 10}
	for _, num := range nums {
		go square(num) //запуск горутины (потока), где square()-потоковая функция
	}
}

func main() {
	do()
	time.Sleep(1000)
	fmt.Println("Итоговая сумма", sum)
}
