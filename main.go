package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader, countLines bool) int {
	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	// If the count lines flag is not set, we want to count words so we define
   // the scanner split type to words (default is split by lines)

    if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	// Defining a counter
	wc := 0

	// For every word scanned, increment the counter
	for scanner.Scan() {
		wc++
	}

	// return the total
	return wc
}

func main() {
	lines := flag.Bool("l", false, "Count lines")

	flag.Parse()


	// Calling the count function to count the number of words (or lines)
// received from the Standard Input and printing it out
	fmt.Println(count(os.Stdin, *lines))
}