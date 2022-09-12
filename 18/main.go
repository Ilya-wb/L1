package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
)

// Создание типа Counter с переменной счяетчика типа int и мьютексом
type Counter struct {
	num int
	sync.Mutex
}

// Метод типа Counter, инкрементирует счетчик
func (c *Counter) Inc() {
	// Захват мьютекса
	c.Lock()
	// Отпускает мьютекс после завершения всех операций в функции
	defer c.Unlock()
	// Инкрементирование счетчика
	c.num += 1
}

// Метод типа Counter, возвращает значение счетчика
func (c *Counter) Value() int {
	return c.num
}

func main() {
	// Объявление переменной типа Counter и инициализация счетчика
	cnt := &Counter{
		num: 0,
	}

	// Создание канала, по сигналу которого будет выводиться значение счетчика в конце работы программы
	finish := make(chan struct{})
	// Запуск горутины с воркером Do()
	go Do(cnt, finish)
	// Конструкция select, ожидающая приема сигнала по каналу finish
	select {
	case <-finish:
		log.Printf("Значение счётчика: %d", cnt.Value())
	}
}

func Do(cnt *Counter, finish chan struct{}) {
	// Запуск ожидания завершения всех горутин
	wg := sync.WaitGroup{}

	// Запуск 10 горутин, выполняющих большое количество арифметических операций
	for i := 0; i < 10; i++ {
		// Инициализация счетчика ожидания
		wg.Add(1)

		// Запуск оцередной горутины в цикле
		go func(num int, cnt *Counter, wg *sync.WaitGroup) {
			// По окончанию работы горутины счетчик wg обнулится
			defer wg.Done()
			// Некоторые арифметические опреации, для нагрузки горутины
			a := rand.Float64()
			b := rand.Float64()
			for j := 0; j < 100; j++ {
				c := a / b
				c = a * b
				fmt.Print(c)
			}
			cnt.Inc()
		}(i, cnt, &wg)
	}
	// Ождидание, пока счетчики не будут равны 0
	wg.Wait()
	// Отправлоение сигнала об окончании работы по каналу finish
	finish <- struct{}{}
	// Закрытие канала
	close(finish)
}
