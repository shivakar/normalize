# normalize

Normalize datasets using various statistics

## Installation

```
go get -u -x github.com/shivakar/normalize
```

### Usage

```
Normalize is a tool to normalize datasets using various statistics.

Usage:
    normalize [options] <input-filename>
  -d string
        delimiter to parse fields in input file (default ",")
  -o string
        filename to write output (default: write to Stdout)
  -out-delimiter string
        delimiter to seperate fields in output file (default ",")
  -trim-whitespace
        trim whitespace around lines and fields before parsing (default true)
```

```
normalize -d "    " -o output.csv input.csv
```
