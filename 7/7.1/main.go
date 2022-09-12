package main

import (
	"fmt"
	"sync"
)

// Возведение числа в квадрат и заполнение m map
func square(num int, m map[int]int, mMutex sync.RWMutex) {
	mMutex.Lock() // Мьютекс в данном случае нужен для того, чтобы избежать ошибки "concurrent map write/read"
	m[num] = num * num
	fmt.Printf("map[%d] = %d\n", num, m[num])
	mMutex.Unlock() // Отпускается мьютекс
}

// В данной функции n раз запускается горутина, где n-количество элементов среза nums
func do() {
	m := make(map[int]int)
	var mMutex sync.RWMutex
	nums := [...]int{1, 2, 3, 4, 5}
	for _, num := range nums {
		go square(num, m, mMutex) //запуск горутины (потока), где square()-потоковая функция
	}
	fmt.Scanln()
}

func main() {
	//m := make(map[int]int)
	do()
}
