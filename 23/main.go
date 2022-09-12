package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	v1(a, 2)
	a = []int{1, 2, 3, 4, 5, 6, 7}
	v2(a, 2)
}

// Не меняется порядок следования элементов
func v1(a []int, i int) {
	// Выполняется сдвиг влево на 1 элемент хвоста среза
	copy(a[i:], a[i+1:])
	// Удаление последнего элемента
	a[len(a)-1] = 0
	// Усечение среза
	a = a[:len(a)-1]
	fmt.Println(a)
}

// Порядок следования элементов меняется
func v2(a []int, i int) {
	// Копирование последнего элемента в индекс i
	a[i] = a[len(a)-1]
	// Удаление последнего элемента
	a[len(a)-1] = 0
	// Усечение среза
	a = a[:len(a)-1]
	fmt.Println(a)
}
