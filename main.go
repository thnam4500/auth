package main

import (
	"fmt"
	"log"
)

func main() {
	store, err := NewPostgresStorage()
	if err != nil {
		log.Fatalln(err)
	}

	err = store.Init()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("hello world")
	server := NewAPIServer(":3000", store)
	server.Run()
}
