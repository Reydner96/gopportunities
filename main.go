package main

import (
	"fmt"

	"github.com/Reydner96/gopportunities/router"
)

func main() {
	// Initialize Configs
	err error := config.Init()
	if err != nil {
		panic(err)
		fmt.Println(err)
		return
	}

	// Initialize Router
	router.Initialize()
}
