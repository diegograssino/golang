package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("Hello Golang")
	fmt.Println("UUID: ", uuid.New().String())
}