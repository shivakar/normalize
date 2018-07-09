package main

import "fmt"

// Normalize normalizes a Dataset using the standard score statistic
func Normalize(ds Dataset) (Dataset, error) {
	out := make(Dataset, len(ds))
	copy(out, ds)

	mu, sd, err := MeanStdDev(out)
	if err != nil {
		return nil, fmt.Errorf("error calculating mean and std. dev: %v", err)
	}

	for i, r := range out {
		for j, v := range r {
			out[i][j] = (v - mu[j]) / sd[j]
		}
	}

	return out, nil
}
