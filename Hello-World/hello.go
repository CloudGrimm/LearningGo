package main

import "fmt"

//Constants should improve performance of your application
const englishHelloPrefix = "Hello, "

func Hello(name string) string { //this returns a string
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Chris"))
}
