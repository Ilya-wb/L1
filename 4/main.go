package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var workerPoolSize int
	fmt.Println("Введите количество воркеров")
	_, err := fmt.Fscan(os.Stdin, &workerPoolSize)
	if err != nil {
		panic(err)
	}
	// создание буферизованного канала sigCh, который принимает сигналы типа os.Signal
	sigCh := make(chan os.Signal, 1)
	// Регистрация канала sigCh для того, чтобы он получал специфичные сигналы, в данном случае системный
	// сигнал остановки программы
	signal.Notify(sigCh, syscall.SIGINT)
	go interruptControl(sigCh)
	run(workerPoolSize)
}

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

func run(workerPoolSize int) {
	// Создание буферизованного канала dataCh с количеством буферов, равным workerPoolSize
	dataCh := make(chan int, workerPoolSize)
	// Горутина, выполняющая заполнение канала данными
	go fillChannel(dataCh)
	// создание WaitGroup
	wg := new(sync.WaitGroup)
	for i := 0; i < workerPoolSize; i++ {
		workerNumber := i
		wg.Add(1) // Инкрементирование счетчика WaitGroup
		go func() {
			defer wg.Done() // Декрементирует счетчик WaitGroup, когда горутина завершает работу
			for data := range dataCh {
				fmt.Printf("worker(%d) data from channel: %d\n", workerNumber, data)
			}
		}()
	}
	// Пока счетчик WaitGroup равен 0 выполняется ожидание
	wg.Wait()
}

func fillChannel(dataChan chan int) {
	//for i := 0; i < 10; i++ {
	for {
		val := rand.Int()
		dataChan <- val
		fmt.Print(" ", val)
		time.Sleep(time.Millisecond * 1)
	}
}
