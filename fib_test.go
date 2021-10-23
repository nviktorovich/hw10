package main

import (
	"testing"
)

type testdata struct {
	number int
	fibValue int
}

var tests = []testdata{{8, 13}, {12, 89}, {23, 17711} }

//проверка корректности результата
func TestFibFunction(t *testing.T) {
	var v int
	for i := range tests{
		v = fib(tests[i].number)
		if v != tests[i].fibValue{
			t.Errorf("\nuncorrect result: number %d expected %d, got %d\n", tests[i].number, tests[i].fibValue, v)
		}
	}
}

func BenchmarkFib10(b *testing.B) {
	// run the fib function b.N times
	for n := 0; n < b.N; n++ {
		fib(10)
	}
}