package singleton

import "fmt"

func RunSingleton() {
	for i := 0; i < 30; i++ {
		go getInstance()
	}

	fmt.Scanln()
}
