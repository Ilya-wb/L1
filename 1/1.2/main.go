package main

import "fmt"

// Все методы типа Temperature доступны через метод Report,
// так как тип Temperature встроен в метод Report, можно сказать,
// что тип Report наследует поля и методы типа Temperature
type Report struct {
	sol         int
	Temperature // В данном случае тип Temperature встроен в report
	Location
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
	report := Report{sol: 15, Temperature: t, Location: bradbury}
	// Доступ к полю типа Temperature через тип Report
	fmt.Printf("a balmy %v° C\n", report.high)
	// Доступ к методу average() из типа Temperature
	fmt.Printf("average from t %vC\n", t.average())
	// Доступ к методу average() осуществляется из типа Report,
	// так как тип Temperature встроен в метод Report, а метод average() пренадлежит типу Temperature
	fmt.Printf("average from report %v° C\n", report.average())
}

// Таким образом метод average() встраивается в структуру t
// Методу становятся доступны поля данной структуры
func (t Temperature) average() Celsius {
	return (t.high + t.low) / 2
}

// Плюсом такой реализации является остутствие необходимости объявлять 2
// метода average() для каждого из типов, как было показано в предыдущей версии
