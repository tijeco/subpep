package main

import (
	"flag"
	"fmt"
	"os"
)

var p = fmt.Println
var help bool
var idField, pepField, seqField int
var inFile, sep, method, outFile string

func init() {
	flag.BoolVar(&help, "help", false, "print usage")
	flag.BoolVar(&help, "h", false, "print usage (shorthand)")

	flag.IntVar(&idField, "id", 0, "header column number")
	flag.IntVar(&pepField, "pep", 1, "header column number")
	flag.IntVar(&seqField, "seq", 2, "header column number")

	flag.StringVar(&inFile, "in", "", "input file")
	flag.StringVar(&outFile, "out", "subpepOut.json", "output file")
	flag.StringVar(&sep, "sep", ",", "input file field delimiter")
	flag.StringVar(&method, "method", "one_per_pep", "subsetting procedure")

	flag.Parse()

	if help || inFile == "" {
		flag.Usage()
		os.Exit(1)
	}

}

func main() {
	p(sep)
}
