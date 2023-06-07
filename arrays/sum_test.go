package arrays

import (
	"reflect"
	"testing"
)


func TestSumAll(t *testing.T) {
	got := SumAll([]int{3,4},[]int{5,5})
	want := []int{7,10}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func SumAll(numbersToSum ...[]int) []int{
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums

}