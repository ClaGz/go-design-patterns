package adapter

import "fmt"

type MacAdapter struct {
	macMachine *Mac
}

func (w *MacAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.macMachine.insertIntoUSBPort()
}
