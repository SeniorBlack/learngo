package main

import "flag"
import "fmt"

var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorte values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")
var help *string = flag.String("h", "", "USAGE: sorter –i <in> –o <out> –a <qsort|bubblesort>")

func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =",
			*algorithm)
	}
}
