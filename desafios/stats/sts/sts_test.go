package sts

import "testing"

type testPair struct {
	values []float64
	sum    float64
	avg    float64
	min    float64
	max    float64
	median float64
}

var tests = []testPair{
	{[]float64{1., 1., 1.}, 3., 1., 1., 1., 1.},
	{[]float64{1.}, 1., 1., 1., 1., 1.},
	{[]float64{1., 2.}, 3., 1.5, 1., 2., 1.5},
	{[]float64{-1., -1.}, 0., 0., -1., -1., -1},
	{[]float64{1., 2., 3.}, 6., 2., 1., 3., 2.},
}

func testSomaInt(t *testing.T) {
	for _, pair := range tests {
		total := SomaFloat64(pair.values)
		if total != pair.sum {
			t.Error(
				"Para: ", pair.values,
				"Esperado: ", pair.sum,
				"Realizado: ", total,
			)
		}

	}
}

func testMediaInt(t *testing.T) {
	for _, pair := range tests {
		total := AvgFloat64(pair.values)
		if total != pair.avg {
			t.Error(
				"Para: ", pair.values,
				"Esperado: ", pair.avg,
				"Realizado: ", total,
			)
		}

	}
}

func testMinInt(t *testing.T) {
	for _, pair := range tests {
		total := MinFloat64(pair.values)
		if total != pair.min {
			t.Error(
				"Para: ", pair.values,
				"Esperado: ", pair.min,
				"Realizado: ", total,
			)
		}

	}
}

func testMaxInt(t *testing.T) {
	for _, pair := range tests {
		total := MaxFloat64(pair.values)
		if total != pair.max {
			t.Error(
				"Para: ", pair.values,
				"Esperado: ", pair.max,
				"Realizado: ", total,
			)
		}

	}
}

func testMedianInt(t *testing.T) {
	for _, pair := range tests {
		total := MedianFloat64(pair.values)
		if total != pair.median {
			t.Error(
				"Para: ", pair.values,
				"Esperado: ", pair.median,
				"Realizado: ", total,
			)
		}

	}
}
