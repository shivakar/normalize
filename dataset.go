package main

import (
	"fmt"
	"math"
)

// Dataset is an 2-d array of floats
type Dataset [][]float64

// Sum returns the sum of a Dataset along the columns (dim=0)
func Sum(ds Dataset) ([]float64, error) {
	return SumAlong(ds, 0)
}

// SumAlong return the sum of a Dataset along the specified dimension
func SumAlong(ds Dataset, dim int) ([]float64, error) {
	if len(ds) < 1 {
		return nil, fmt.Errorf("received empty dataset")
	}
	oSize := 0
	switch dim {
	case 0:
		oSize = len(ds[0])
	case 1:
		oSize = len(ds)
	default:
		return nil, fmt.Errorf("not implemented for dim > 1")
	}
	out := make([]float64, oSize)

	for i, r := range ds {
		for j, v := range r {
			if dim == 0 {
				out[j] += v
			} else {
				out[i] += v
			}
		}
	}

	return out, nil
}

// Mean returns the mean of a Dataset along the columns (dim = 0)
func Mean(ds Dataset) ([]float64, error) {
	return MeanAlong(ds, 0)
}

// MeanAlong returns the mean of the dataset along the specified dimension
func MeanAlong(ds Dataset, dim int) ([]float64, error) {
	t, err := SumAlong(ds, dim)
	if err != nil {
		return nil, fmt.Errorf("error calculating sum: %v", err)
	}
	for i := range t {
		if dim == 0 {
			t[i] /= float64(len(ds))
		} else {
			t[i] /= float64(len(ds[0]))
		}
	}
	return t, nil
}

// MeanStdDev returns the mean and population standard deviation of the dataset
func MeanStdDev(ds Dataset) ([]float64, []float64, error) {
	mean, err := Mean(ds)
	if err != nil {
		return nil, nil, fmt.Errorf("error computing mean: %v", err)
	}
	variance := make([]float64, len(mean))
	for _, r := range ds {
		for i, v := range r {
			variance[i] += (v - mean[i]) * (v - mean[i])
		}
	}

	for i := range variance {
		variance[i] /= float64(len(ds))
		variance[i] = math.Sqrt(variance[i])
	}

	return mean, variance, nil
}
