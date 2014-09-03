package roulette

import (
	"math"
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

func round(n float64) int {
	return int(math.Floor(n + 0.5))
}

func percent(part int, whole int) int {
	return round(float64(part) / float64(whole) * 100)
}

func around(n int, ref int) bool {
	return (n >= (ref-1) && n <= (ref+1))
}

func TestProbabilityDistribution(t *testing.T) {
	r := NewRoulette()
	r.Add("A", 1)
	r.Add("B", 1)
	r.Add("C", 2)

	runs := 10000
	results := make(map[string]int)
	for i := 0; i < runs; i++ {
		results[r.Roll().(string)] += 1
	}

	pA := percent(results["A"], runs)
	pC := percent(results["C"], runs)
	if !around(pA, 25) {
		t.Fatalf("Roll results are not weighted accordingly. A appears %d%% rather than %d%%", pA, 5)
	}
	if !around(pC, 50) {
		t.Fatalf("Roll results are not weighted accordingly. C appears %d%% rather than %d%%", pC, 50)
	}

}

func TestProbabilityDistributionMinimum(t *testing.T) {
	r := NewRoulette()
	r.LessIsBetter()
	r.Add("A", 1)
	r.Add("B", 1)
	r.Add("C", 2)

	runs := 10000
	results := make(map[string]int)
	for i := 0; i < runs; i++ {
		results[r.Roll().(string)] += 1
	}

	pA := percent(results["A"], runs)
	pC := percent(results["C"], runs)
	if !around(pA, 40) {
		t.Fatalf("Roll results are not weighted accordingly. A appears %d%% rather than %d%%", pA, 40)
	}
	if !around(pC, 20) {
		t.Fatalf("Roll results are not weighted accordingly. C appears %d%% rather than %d%%", pC, 20)
	}

}
