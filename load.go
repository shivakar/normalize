package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"strconv"
	"strings"
)

var (
	trimws = flag.Bool("trim-whitespace", true, "trim whitespace around lines and fields before parsing")
)

// LoadDataset loads contents of a reader into memory parsing it as a dataset using the supplied delimiter
func LoadDataset(f io.Reader, delimiter string) (Dataset, error) {
	out := Dataset{}

	nFields := 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := string(s.Text())
		if *trimws {
			line = strings.TrimSpace(line)
		}
		fields := strings.Split(line, delimiter)
		if len(fields) == 1 && fields[0] == "" {
			return nil, fmt.Errorf("empty lines in the input file are not supported")
		}
		if nFields == 0 {
			nFields = len(fields)
		} else if len(fields) != nFields {
			return nil, fmt.Errorf("number of fields in record not consistent")
		}
		record := make([]float64, nFields)
		for i, v := range fields {
			if *trimws {
				v = strings.TrimSpace(v)
			}
			vf, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return nil, fmt.Errorf("error parsing field: %v", err)
			}
			record[i] = vf
		}
		out = append(out, record)
	}
	return out, nil
}
