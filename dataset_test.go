package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	ds := Dataset{{1, 2}, {3, 4}, {5, 6}}
	expected0 := []float64{9, 12}
	expected1 := []float64{3, 7, 11}

	a0, err := Sum(ds)
	assert.NoError(t, err)
	assert.Equal(t, expected0, a0)

	a0, err = SumAlong(ds, 0)
	assert.NoError(t, err)
	assert.Equal(t, expected0, a0)

	a1, err := SumAlong(ds, 1)
	assert.NoError(t, err)
	assert.Equal(t, expected1, a1)

	e4, err := SumAlong(ds, 4)
	assert.Error(t, err)
	assert.Nil(t, e4)

	ds = Dataset{}
	e0, err := SumAlong(ds, 1)
	assert.Error(t, err)
	assert.Nil(t, e0)
}

func TestMean(t *testing.T) {
	ds := Dataset{{1, 2}, {3, 4}, {5, 6}}
	emu0 := []float64{3, 4}
	emu1 := []float64{1.5, 3.5, 5.5}

	mu0, err := Mean(ds)
	assert.NoError(t, err)
	assert.Equal(t, emu0, mu0)

	mu1, err := MeanAlong(ds, 1)
	assert.NoError(t, err)
	assert.Equal(t, emu1, mu1)
}

func TestMeanStdDev(t *testing.T) {
	ds := Dataset{{1, 2}, {3, 4}, {5, 6}}
	emu := []float64{3, 4}
	esd := []float64{1.632993161855452, 1.632993161855452}

	mu, sd, err := MeanStdDev(ds)
	assert.NoError(t, err)
	assert.Equal(t, emu, mu)
	assert.Equal(t, esd, sd)
}
