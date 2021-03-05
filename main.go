package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var p = fmt.Println
var help, skipHeader bool
var leftPep, idField, pepField, seqField int
var inFile, method, sep, outFile string

// var sep rune
var seed int64

func init() {
	flag.BoolVar(&help, "help", false, "print usage")
	flag.BoolVar(&help, "h", false, "print usage (shorthand)")
	flag.BoolVar(&skipHeader, "skipHeader", false, "input has header column")

	flag.IntVar(&leftPep, "leftPep", 0, "if no peptide column present, use this to split the sequence column into pep column of k length")
	flag.IntVar(&idField, "id", 0, "header column number")
	flag.IntVar(&seqField, "seq", 1, "header column number")
	flag.IntVar(&pepField, "pep", 2, "header column number")
	flag.Int64Var(&seed, "seed", time.Now().UnixNano(), "seed to use for randomness")

	flag.StringVar(&inFile, "in", "", "input file (required)")
	flag.StringVar(&outFile, "out", "subpepOut.json", "output file")
	flag.StringVar(&sep, "sep", ",", "input file field delimiter")
	flag.StringVar(&method, "method", "one_per_pep", "subsetting procedure")

	flag.Parse()

	if help || inFile == "" {
		flag.Usage()
		os.Exit(1)
	}

}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = rune(sep[0])
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	} else if skipHeader {
		records = records[1:]
	}

	return records
}

func randSubseq(seq string, l int, seed int64) (string, error) {
	var subseq string
	rand.Seed(seed)
	if len(seq) > l {
		fmt.Println(len(seq) - l)
		start := rand.Intn(len(seq) - l)
		subseq = seq[start : start+l]
	} else {
		return "", errors.New("Sequence shorter than requested subsequence")
	}
	return subseq, nil
}

func makeLeftPep(csvRecords [][]string, k, idField, seqField int) map[string][]string {
	outDict := map[string][]string{
		"id":     []string{},
		"pep":    []string{},
		"subseq": []string{}}
	for _, row := range csvRecords {
		seq := row[seqField]
		id := row[idField]
		if len(seq)+k >= 2*k {
			pep := seq[:k]
			remainingSeq := seq[k:]
			subseq, err := randSubseq(remainingSeq, k, seed)
			if err == nil {
				outDict["id"] = append(outDict["id"], id)
				outDict["pep"] = append(outDict["pep"], pep)
				outDict["subseq"] = append(outDict["subseq"], subseq)
			}

		}

	}
	return outDict
}

func main() {
	p(inFile)
	csvTest := readCsvFile(inFile)
	p(csvTest)
	if leftPep == 0 {

	} else {
		testOut := makeLeftPep(csvTest, leftPep, idField, seqField)
		p(testOut)
	}

}
