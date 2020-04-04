package main

func main() {
	f := "fizz"
	b := "buzz"
	for i := 1; i < 21; i++ {
		if i%3 == 0 {
			println(f)
		} else if i%5 == 0 {
			println(b)
		} else if i%5 == 0 && i%3 == 0 {
			println(f + b)
		} else {
			println(i)
		}
	}
}
