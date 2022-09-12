package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Создание двух каналов с1-передатчик начального числа, с2-передатчик квадрата числа
	c1 := make(chan int)
	c2 := make(chan int)

	var arr [1000]int

	// Заполнение массива значениями int от 1 до 10
	for i, _ := range arr {
		arr[i] = rand.Intn(10)
		fmt.Print(arr[i])
	}

	go c0(c1, arr)
	go conv1(c1, c2)
	go conv2(c2)

	fmt.Scanln()
}

// Горутина последовательно берет каждый элемент массива и отправляет его по каналу с1
// Также создается задержка в 1 секунду для удобства оценки результата работы программы в консоле
func c0(c1 chan int, arr [1000]int) {
	for _, val := range arr {
		fmt.Printf("c0: %d, ", val)
		c1 <- val
		time.Sleep(time.Second * 1)
	}
}

// Данная горутина принимает значение из канала с1, возводит его в квадрат и отправляет по каналу с2
func conv1(c1 chan int, c2 chan int) {
	for {
		val := <-c1
		fmt.Printf("conv1: %d, ", val)
		val = val * val
		fmt.Printf("val^2: %d, ", val)
		c2 <- val
	}
}

// Данная горутина принимает квадрат числа из канала с2 и выводит его на экран
func conv2(c2 chan int) {
	for {
		fmt.Printf("conv2: %d\n", <-c2)
	}
}
