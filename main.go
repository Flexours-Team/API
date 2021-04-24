package main

import (
	"fmt"
	"github.com/Flexours-Team/API/internal/api"
	"os"
)

func main() {
	server, err := api.NewServer()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	server.Start()
}
