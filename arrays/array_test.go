package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("Array of 6 numbers",func(t *testing.T){

	numbers := []int{1,2,3,4,5,6}

	got := Sum(numbers)
	want := 21
	
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
})
	t.Run("Slice: Collection of any size",func(t *testing.T) {
		numbers := []int{10,20,30,40}

		got := Sum(numbers)
		want := 100

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}

func Sum(numbers []int)int {
	sum := 0
	for _ , number := range numbers {
		sum += number
	}
	return sum

}

