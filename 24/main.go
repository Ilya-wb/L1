package main

import (
	"fmt"
	"math"
)

// Создается тип данных для хранения координат одной точки
type Cords struct {
	x, y float64
}

func main() {
	// Инициализация координат двух точек
	a := Cords{3, 3}
	b := Cords{6, 6}
	// Рассчет расстояния по формуле
	d := math.Sqrt((math.Pow((b.x-a.x), 2) + math.Pow((b.y-a.y), 2)))
	fmt.Println("Расстояние между а и b", d)
}
