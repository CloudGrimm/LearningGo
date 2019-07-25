//packages are a way of grouping up related Go code together
package main

//With import "fmt" we are importing a package which contains the Println function that we use to print
import "fmt"

//Constants should improve performance of your application
const spanish = "Spanish"
const english = "English"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

//The func keyword is how you define a function with a name and a body
func Hello(name, language string) string { //this returns a string
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}
func greetingPrefix(language string) (prefix string) {

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Chris", english))
}
