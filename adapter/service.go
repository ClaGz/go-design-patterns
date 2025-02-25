package adapter

import "fmt"

type Mac struct {
}

type Windows struct {
}

func (m *Mac) insertIntoUSBPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

func (m *Windows) insertIntoUSBPort() {
	fmt.Println("Lightning connector is plugged into windows machine.")
}
