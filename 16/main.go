package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var arr1 []int
	for i := 0; i < 25; i++ {
		arr1 = append(arr1, rand.Intn(25))
	}
	fmt.Printf("До сортировки: ")
	fmt.Println(arr1)
	// Данная функция библиотеки sort реализована с использованием быстрой версии quicksort
	sort.Ints(arr1)

	fmt.Printf("sort.Ints(): После сортировки: ")
	fmt.Println(arr1)

	var arr2 []int
	for i := 0; i < 25; i++ {
		arr2 = append(arr2, rand.Intn(25))
	}
	fmt.Printf("До сортировки: ")
	fmt.Println(arr2)

	fmt.Println("quicksort: После сортировки: ", quicksort(arr2))
}

// Данная функция представляет из себя классическую быструю сортирову
func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}
