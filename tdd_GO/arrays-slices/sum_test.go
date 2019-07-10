package main

import (
	"reflect"
	"testing"
)

/**
The syntax of slices and arrays is similar
where the only difference is ommitting the
size when declaring them
**/

/**
Never forget test code in the refactoring
attempts
**/

/**
It is important to question the value of your tests.
It should not be a goal to have as many tests as possible,
but rather to have as much confidence as possible in your code base.
Having too many tests can turn in to a real problem and it just adds
more overhead in maintenance.
Every test has a cost.
*/

/**
TEST COVERAGE
Go's built-in testing toolkit features a coverage tool,
which can help identify areas of your code you have not covered.
I do want to stress that having 100% coverage should not be your goal,
it's just a tool to give you an idea of your coverage. If you have been strict with TDD,
it's quite likely you'll have close to 100% coverage anyway.
Try running
go test -cover
**/

/**
Compile time errors are our friend because they help us write software that works,
runtime errors are our enemies because they affect our users.
**/

// /**Original test***/
func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}

	got := SumAll(slice1, slice2)
	want := []int{6, 6}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v giving %v and %v, ", got, want, slice1, slice2)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}
