package ds

import "fmt"

func fibonacci(n int) string {
	a, b := 0, 1
	output := fmt.Sprintf("%d %d", a, b)
	for i := 2; i < n; i++ {
		c := a + b
		output += fmt.Sprintf(" %d", c)
		a = b
		b = c
	}
	return output
}
