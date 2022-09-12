package main

func main() {
	client := &Client{}
	mac := &Mac{}
	// В данном случае клиент подключается к mac.os напрямую, без адаптера
	client.insertSquareUsbInComputer(mac)
	// В этом случае клиент подключается через адаптер
	// Функция insertSquareUsbInComputer печатает Insert circle port into windows, так как она метод типа WindowsAdapter
	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}
	// Вместо машины передается адаптер, который в свою очередь вызывает insertInCirclePort который печатает
	// Insert circle port into windows machine
	client.insertSquareUsbInComputer(windowsMachineAdapter)
}
