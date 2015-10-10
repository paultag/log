package main

import (
	"time"
)

var root = "/home/paultag/log/"

func main() {
	when := time.Now()

	if err := Create(root, when); err != nil {
		panic(err)
	}

	if err := Log(root, when, "Hello, World"); err != nil {
		panic(err)
	}
}
