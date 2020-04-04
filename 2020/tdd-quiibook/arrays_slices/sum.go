package arrays_slices

func Sum(numbers []int) int{
	var sum int
	//for i :=0; i < 5; i++{
	//	sum += numbers[i]
	//}
	/**
	range lets you iterate over an array. Every time it
	is called it returns two values, the index and
	the value. We are choosing to ignore the index value by
	using _ blank identifier.
	**/
	for _, number := range numbers{
		sum += number
	}
	return sum
}
