package main

import (
	"fmt"

	"github.com/mixcleiton/base/pkg/builder"
	"github.com/mixcleiton/base/pkg/factory"
	"github.com/mixcleiton/base/pkg/singleton"
)

func main() {
	fmt.Println("Using Build Pattern")
	builder.RunBuilder()

	fmt.Println("")
	fmt.Println("")
	fmt.Println("Using Run Pattern")
	factory.RunFactory()

	fmt.Println("")
	fmt.Println("")
	fmt.Println("Using Run Pattern")
	singleton.RunSingleton()
}
