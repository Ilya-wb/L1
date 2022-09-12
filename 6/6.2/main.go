package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Данная программа завершается с нажатием Ctrl+C, также при этом завершается и горутина
func interruptControl(sigCh chan os.Signal) {
	// Данный вечный цикл постоянно слушает stdin и в случае нажатия клавиш Ctrl+C программа моментально завершится
	// спомощью команды os.Exit(0), при этом exit code будет 0, соответственно. Это штатный способ
	// завершения программы. Спомощью завершения программы без использования os.Exit(0) код будет случайным числом
	for {
		s := <-sigCh
		switch s {
		case syscall.SIGINT:
			fmt.Printf("Get ~C\n")
			os.Exit(0)
		default:
			fmt.Printf("Unknown signal\n")
		}

	}
}

func main() {
	var i int
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT)
	go interruptControl(sigCh)
	go func() {
		for {
			fmt.Println(i)
			i++
		}
	}()
	fmt.Scanln()
}
