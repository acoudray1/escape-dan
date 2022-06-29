package main

import (
	"fmt",
	"time"
)

func main() {
	for {
		fmt.Println("Hello wrld!!")
		time.Sleep(time.Second * 3)
	}
}