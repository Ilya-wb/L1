package main

import (
	"fmt"
	"sort"
)

func main() {
	sl := []int{1, 2, 3, 6, 9, 15, 16, 19, 20}
	//Данная функция реализована на основе алгоритма бинарного поиска, сложность O(log n)
	idx1 := sort.SearchInts(sl, 6)
	fmt.Println(idx1)
	// Данная функция представляет из себя общий бинарный поиск, первый аргумент-длина массива
	// второй-функция, которая возвращает false, если текущий элменет больше либо равен искомому
	idx2 := sort.Search(len(sl), func(i int) bool { return 6 <= sl[i] })
	fmt.Println(idx2)
}
