package main

import "fmt"

func thread() {
	const HelloWorld = "Hello World"
	fmt.Println("I am a thread")

}

func main() {
	go thread()
	fmt.Println("Hello World!")
}
