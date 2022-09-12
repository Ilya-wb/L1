package main

import (
	"fmt"
	"sync"
)

// В данной реализации используется syncmap
// Syncmap используется в основном в тех случаях, когда количество чтения map на много больше чем количество записей
// Если в решении много записей, то из-за syncmap потеряется скорость
// Syncmap гарантирует константное время доступа к map вне зависимости от количества параллельных чтений
// Возведение числа в квадрат и заполнение m sync.Map
func square(num int, m sync.Map) {
	m.Store(num, num*num)
	val, _ := m.Load(num)
	fmt.Printf("map[%d] = %d\n", num, val)
}

// В данной функции n раз запускается горутина, где n-количество элементов среза nums
func do() {
	var m = sync.Map{}
	nums := [...]int{1, 2, 3, 4, 5}
	for _, num := range nums {
		go square(num, m) //запуск горутины (потока), где square()-потоковая функция (worker)
	}
	fmt.Scanln()
}

func main() {
	do()
}
