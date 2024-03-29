package roulette

import (
	"math/rand"
	"time"
)

type Roulette interface {
	Add(item interface{}, weight float64)
	Roll() interface{}
	LessIsBetter()
}

type slice struct {
	position float64
	item     interface{}
}

type roulette struct {
	slices  []slice
	total   float64
	rng     func() float64
	minimum bool
}

func (r *roulette) LessIsBetter() {
	r.minimum = true
}

func (r *roulette) Add(item interface{}, weight float64) {
	if r.minimum {
		r.total += 1 / (weight + 0.0000000000000001)
	} else {
		r.total += weight
	}

	r.slices = append(r.slices, slice{item: item, position: r.total})
}

func (r *roulette) Roll() interface{} {
	p := r.rng() * r.total
	var j int
	for i, slice := range r.slices {
		if p < slice.position {
			j = i
			break
		}
	}
	return r.slices[j].item
}

func rnd() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

var defaultNrg = func() float64 {
	return rnd().Float64()
}

func NewRoulette() Roulette {
	return NewRouletteRng(defaultNrg)
}

func NewRouletteRng(f func() float64) Roulette {
	return &roulette{rng: f}
}
