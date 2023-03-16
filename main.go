// Program to illustrate multiple goroutines

package main

import (
	"fmt"
	"time"
)

func display(message string) {

	fmt.Println(message)

}

func main() {

	// run two different goroutine
	go display("Fruit 1")
	go display("Fruit 2")
	go display("Fruit 3")

	// to sleep main goroutine for 1 sec
	time.Sleep(time.Second * 1)
}
