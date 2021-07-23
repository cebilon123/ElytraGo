package main

import (
	"fmt"
	"github.com/cebilon123/ElytraGo/server"
)

func main() {
	fmt.Printf("Starting ElytraGO on port: 8084")

	server.New().Create().Start()
}
