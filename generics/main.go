package main

import "fmt"

type Number interface {
	int64 | float64
}

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

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
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

	fmt.Printf("Non-Generics Sums: %v and %v\n",
		SumInts(ints), SumFloats(floats))
	fmt.Printf("Generics Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints), SumIntsOrFloats[string, float64](floats))
	fmt.Printf("Generic Sum with Consistraint: %v and %v\n",
		SumNumbers(ints), SumNumbers(floats))
}
