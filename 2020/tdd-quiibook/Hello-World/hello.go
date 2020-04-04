package main // Packages are ways of grouping up related Go code together

import "fmt"

const englishHelloPrefix  = "Hello, "
func Hello(name string) string{
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name

}

func main(){//func defines a function with a name and a body
	fmt.Println(Hello("Tinashe"))
}