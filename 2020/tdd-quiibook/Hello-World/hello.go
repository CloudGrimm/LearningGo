package main // Packages are ways of grouping up related Go code together

import "fmt"

func Hello(name string) string{
	return "Hello, " + name
}
func main(){//func defines a function with a name and a body
	fmt.Println(Hello("Tinashe"))
}