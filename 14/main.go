package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i interface{} // Пустой интерфейс может содержать значение любого типа
	//describe(i)

	i = 42
	x := reflect.TypeOf(i)         // В данном случае тип сохранится в переменной x типа reflect.Type
	fmt.Println(x)                 // Выведет int
	fmt.Printf("(%v, %T)\n", i, i) // Таким образом можно вывести значение и тип
	// %v - универсальный спецификатор формата, %T - вывод типа данных

	i = "hello"
	x = reflect.TypeOf(i) // То же самое для string
	fmt.Println(x)        // Выведет string
	fmt.Printf("(%v, %T)\n", i, i)

	var c chan int
	i = c
	x = reflect.TypeOf(i) // То же самое для chan int
	fmt.Println(x)        // Выведет int
	fmt.Printf("(%v, %T)\n", i, i)

	var b bool
	i = b
	x = reflect.TypeOf(i) // То же самое для bool
	fmt.Println(x)        // Выведет bool
	fmt.Printf("(%v, %T)\n", i, i)
}
