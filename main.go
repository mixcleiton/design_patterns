package main

import (
	"fmt"

	"github.com/mixcleiton/base/pkg/builder"
	"github.com/mixcleiton/base/pkg/singleton"
)

func main() {
	builder.RunBuilder()

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	singleton.RunSingleton()
}
