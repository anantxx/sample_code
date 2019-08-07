package fib

import (
	"testing"
)

/*var fibStruct = []struct {
	n, expected int
}{
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
	{8, 21},
	{9, 34},
	{10, 55},
	{11, 89},
}

func TestFibData(t *testing.T) {
	for _, tt := range fibStruct {
		if actual := Fib(tt.n); actual != tt.expected {
			t.Errorf("Fib(%d) : expected %d, Actaul %d", tt.n, tt.expected, actual)
		}
	}
}*/

func benchmarkFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(i)
	}
}

func BenchmarkFib1(b *testing.B) { benchmarkFib(1, b) }

/*func BenchmarkFib2(b *testing.B)  { benchmarkFib(2, b) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(3, b) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }
*/
/* func BenchmarkFibWrong(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(n)
	}
} */

/*func BenchmarkFibWrong2(b *testing.B) {
	Fib(b.N)
}*/
