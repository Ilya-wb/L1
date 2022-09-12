package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(1)
	sleep(5 * time.Second)
	fmt.Println(2)
}

func sleep(d time.Duration) {
	// Ждет указанное время, а затем отправляет текущее время по каналу типа Time
	<-time.After(d)
}
