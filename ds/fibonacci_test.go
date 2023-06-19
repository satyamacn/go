package ds

import "testing"
import "fmt"

	func TestFibonacci(t *testing.T) {
		testCases := []struct {
			n    int
			want string
		}{
			{2, "0 1"},
			{5, "0 1 1 2 3"},
			{10, "0 1 1 2 3 5 8 13 21 34"},
		}
	
		for _, tc := range testCases {
			t.Run(fmt.Sprintf("Fibonacci(%d)", tc.n), func(t *testing.T) {
				got := fibonacci(tc.n)
				fmt.Println(got)
				if got != tc.want {
					t.Errorf("Fibonacci(%d) = %q, want %q", tc.n, got, tc.want)
				}
			})
		}
		
	}

	