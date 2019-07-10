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

// func SumAll(numbersToSum ...[]int) []int {
// 	//get the length of the numbers to be summed
// 	lengthOfNumbers := len(numbersToSum)
// 	//create the slice with the number of sums to be calculated
// 	sums := make([]int, lengthOfNumbers)
// 	//iterate from the above
// 	for i, numbers := range numbersToSum {
// 		sums[i] = Sum(numbers)
// 	}
// 	return sums
// }

//refactor

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

/**
This function calculates the totals of the "tails" of each slice
The tail of a collection is all the items apart from the first one "the head"
*/
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}

// func main (){
// 	println(SumAllTails([]int{}))
// }
