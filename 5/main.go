package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// Создание канала по которому будет посылаться ссообщение
	c := make(chan int)
	// Засекаем время начала работы программы
	timeBefore := time.Now().Second()
	// Запускаем горутину, которая завершит работу программы через N секунд
	go interruptControl(timeBefore)
	// Запускаем горутину, которая посылает сообщения
	go sender(c)
	// Запускаем горутину, которая принимает сообщения
	go receiver(c)
	// Задержка, чтобы функционал программы успел отработать до завершения программы
	fmt.Scanln()
}

func interruptControl(start int) {
	for {
		// Рассчет времени, прошедшего с начала запуска в секундах
		diffTime := time.Now().Second() - start
		if diffTime > 5 {
			fmt.Println("exit")
			// Остановка процесса с соответствующем exit code
			os.Exit(0)
		}
	}
}

func sender(c chan int) {
	for {
		// Данные для отправки представляют из себя число типа int и генерируются случайным образом
		i := rand.Int()
		fmt.Println("sended", i)
		// Посылка сгенерированного ранее числа в канал
		c <- i
	}
}

func receiver(c chan int) {
	for {
		// Прием данных из канала
		val := <-c
		fmt.Println("received", val)
		// Задержка, для того, чтобы можно было рассмотреть результат в консоле
		time.Sleep(time.Second * 1)
	}
}
