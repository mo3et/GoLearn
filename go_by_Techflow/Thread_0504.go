package main

import "fmt"

func thread() {
	fmt.Println("I am a thread")

}

func main() {
	go thread()
	fmt.Println("Hello World!")
}
