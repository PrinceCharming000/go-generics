package main

import "fmt"

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, i := range m {
		s += i
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, i := range m {
		s += i
	}
	return s
}

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  32.45,
		"second": 54.41,
	}

	fmt.Printf("Non-Generics Sums: %v and %v\n", SumInts(ints), SumFloats(floats))
}
