package main

import (
	"go-design-patterns/adapter"
	chainofresponsability "go-design-patterns/chain-of-responsability"
	"go-design-patterns/singleton"
)

func main() {
	adapter.Run(adapter.TypeWindows)
	adapter.Run(adapter.TypeMac)

	singleton.Run()

	chainofresponsability.Run()
}
