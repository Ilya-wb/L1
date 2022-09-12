package main

type WindowsAdapter struct {
	windowMachine *Windows
}

func (w *WindowsAdapter) insertInSquarePort() {
	w.windowMachine.insertInCirclePort()
}
