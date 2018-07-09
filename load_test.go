package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadDataset(t *testing.T) {
	data := "1,2\n2,3\n4,5\n"
	expected := Dataset{{1, 2}, {2, 3}, {4, 5}}
	ds, err := LoadDataset(strings.NewReader(data), ",")
	assert.NoError(t, err)
	assert.Equal(t, expected, ds)

	data = "\n\n\n"
	ds, err = LoadDataset(strings.NewReader(data), ",")
	assert.Error(t, err)
	assert.Nil(t, ds)

	data = "1,2\n1,2,3\n"
	ds, err = LoadDataset(strings.NewReader(data), ",")
	assert.Error(t, err)
	assert.Nil(t, ds)

	data = "1,2\n1,a\n"
	ds, err = LoadDataset(strings.NewReader(data), ",")
	assert.Error(t, err)
	assert.Nil(t, ds)
}
