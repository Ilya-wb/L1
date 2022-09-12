package main

import "fmt"

type Report struct {
	sol         int
	temperature Temperature // Поле temperature является структурой типа Temperature
	location    Location    // Поле location является структурой типа Location
}

type Temperature struct {
	high, low Celsius // Поля high и low являются структурами типа Celsius
}

type Location struct {
	lat, long float64
}

// Вместо того, чтобы вдальнейшем не писать float64 объявляется структура Celsius
type Celsius float64

func main() {
	// Объявление и инициализация структур
	bradbury := Location{-4.5895, 137.4417}
	t := Temperature{high: -1.0, low: -78.0}
	// Структуре передаются инициализированные структуры t и bradbury
	report := Report{sol: 15, temperature: t, location: bradbury}

	// Доступ к полю структуры t, которое пренадлежит структуре report
	fmt.Printf("a balmy %v° C\n", report.temperature.high)
	fmt.Printf("average %v° C\n", report.temperature.average())
}

// Таким образом метод average() встраивается в структуру t
// Методу становятся доступны поля данной структуры
func (t Temperature) average() Celsius {
	return (t.high + t.low) / 2
}

func (r Report) average() Celsius {
	return r.temperature.average()
}
