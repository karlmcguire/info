package main

import (
	"fmt"
	"math"
)

// (Four-event source) Consider a source with four events having probabilities
// 1/5, 1/5, 1/5, 2/5.
//
// a) What is the information in bits conveyed by a report that the first event
//    occured?
// b) What is the entropy of the source?

func main() {
	var (
		source = []float64{
			float64(1) / float64(5),
			float64(1) / float64(5),
			float64(1) / float64(5),
			float64(2) / float64(5),
		}
		total = float64(0)
	)

	for _, p := range source {
		total += p * math.Log2(1/p)
	}

	fmt.Printf(
		"a) %f\nb) %f\n",
		total,
		math.Log2(1/source[0]),
	)
}
