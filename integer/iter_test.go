package integer

import "testing"

const repeatCount=5

func TestRepeat(t *testing.T){
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected,repeated)
	}

}

func Repeat(s string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += s
	}
	return repeated
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}