package singleton

import (
	"fmt"
	"sync"
)

type Singleton struct {
}

var lock = &sync.Mutex{}
var instance *Singleton

func getInstance() *Singleton {

	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		if instance == nil {
			fmt.Println("Creating single instance now.")
			instance = &Singleton{}
		} else {
			fmt.Println("Single instance already created.")
		}

	} else {
		fmt.Println("Single instance already created.")
	}

	return instance
}
