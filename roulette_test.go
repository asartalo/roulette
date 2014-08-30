package roulette

import (
	"testing"
)

func TestSimpleExample(t *testing.T) {
	r := NewRoulette()
	r.Add("A", 1)
	if r.Roll() != "A" {
		t.Fatal("Should return same element for single-slice roulette.")
	}
}

func TestCorrectness(t *testing.T) {
	i := 0
	rngResults := []float64{0.51, 0.1, 0.26}
	rng := func() float64 {
		ret := rngResults[i%len(rngResults)]
		i++
		return ret
	}

	r := NewRouletteRng(rng)
	r.Add("A", 1)
	r.Add("B", 1)
	r.Add("C", 2)
	result1 := r.Roll()
	result2 := r.Roll()
	result3 := r.Roll()
	if result2 != "A" && result3 != "B" && result1 != "C" {
		t.Fatal("Roulette is not using number generator properly.")
	}
}
