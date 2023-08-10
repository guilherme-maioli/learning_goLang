package sts

import "testing"

type testPair struct {
	values []float64
	sum    float64
	avg    float64
}

var tests = []testPair{
	{[]float64{1., 1., 1.}, 3., 1.},
	{[]float64{1.}, 1., 1.},
	{[]float64{1., 2.}, 3., 1.5},
	{[]float64{-1., -1.}, 0., 0.},
}

func testSomaInt(t *testing.T) {
	for _, pair := range tests {
		total := SomaFloat64(pair.values)
		if total != pair.sum {
			t.Error("Para: ", pair.values, "Esperado: ", pair.sum, "Realizado: ", total)
		}

	}
}

func testMediaInt(t *testing.T) {
	for _, pair := range tests {
		total := AvgFloat64(pair.values)
		if total != pair.avg {
			t.Error("Para: ", pair.values, "Esperado: ", pair.avg, "Realizado: ", total)
		}

	}
}
