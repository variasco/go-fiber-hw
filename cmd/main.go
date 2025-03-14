package main

import (
	"fmt"
	"variasco/go-fiber-hw/config"
)

func main() {
	conf := config.LoadConfig()
	fmt.Println("Loaded configuration: ", conf)
}
