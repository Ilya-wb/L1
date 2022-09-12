package main

import "fmt"

func selection(quit chan int) {
	for {
		select {
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	quit := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
		quit <- 0
	}()
	selection(quit)
}
