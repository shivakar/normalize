package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	outFilename  = flag.String("o", "", "filename to write output (default: write to Stdout)")
	inDelimiter  = flag.String("d", ",", "delimiter to parse fields in input file")
	outDelimiter = flag.String("out-delimiter", ",", "delimiter to seperate fields in output file")
)

// usage prints the usage of the program
func usage() {
	fmt.Fprintf(os.Stderr, "\nNormalize is a tool to normalize datasets using various statistics.\n\n")
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "    normalize [options] <input-filename>\n")
	flag.PrintDefaults()
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

}

func main() {
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "Invalid usage. Input file not specified.\n")
		flag.Usage()
		os.Exit(2)
	}
	infileName := flag.Arg(0)
	f, err := os.Open(infileName)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("Input file does not exist")
		}
		log.Fatalf("Error opening input file: %v", err)
	}
	defer f.Close()

	var of *os.File
	if *outFilename == "" {
		of = os.Stdout
	} else {
		of, err = os.Create(*outFilename)
		if err != nil {
			log.Fatalf("error creating output file: %v", err)
		}
		defer of.Close()
	}

	ds, err := LoadDataset(f, *inDelimiter)
	if err != nil {
		log.Fatalf("error loading dataset: %v", err)
	}
	ds, err = Normalize(ds)
	if err != nil {
		log.Fatalf("error normalizing dataset: %v", err)
	}

	for _, r := range ds {
		rs := make([]string, 0, len(r))
		for _, v := range r {
			rs = append(rs, fmt.Sprint(v))
		}
		of.WriteString(strings.Join(rs, *outDelimiter) + "\n")
	}

	log.Printf("Done writing outfile: %v", *outFilename)
}
