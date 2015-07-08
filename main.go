package main

import (
	"fmt"

	"github.com/slok/daton/configuration"
)

func main() {
	configuration.LoadSettingsFromFile()
	fmt.Println("Hello world")
}
