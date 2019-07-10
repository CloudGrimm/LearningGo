package main

//Original function
// func Sum(numbers [5]int)int{
// 	sum :=0
// 	for i:=0; i < len(numbers); i++{
// 		sum += numbers[i]
// 	}
// 	return sum
// }
//refactor
func Sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
