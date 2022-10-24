// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) // create counts, a map of string to int values
	files := os.Args[1:]           // files, which is the value of os.Args from 1 and beyond. 0 is the command itself.
	if len(files) == 0 {           // if there are no files
		countLines(os.Stdin, counts) // pass std input into countLines function
	} else {
		for _, arg := range files { // however if there are files, go through them
			f, err := os.Open(arg) // f is the file and the err is any error
			if err != nil {        // If there's an error, print it out
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts) // pass the file into the countLines function
			f.Close()             //close the file
		} // this repeats for each file
	}
	for line, n := range counts { // from the counts file
		if n > 1 { // anything that was tracked, print it out below
			fmt.Printf("%d\t%s\n", n, line)

		}
	}
}

func countLines(f *os.File, counts map[string]int) { // create func, takes f, file pointer and counts map
	input := bufio.NewScanner(f) // Create a new scanner object, input, and pass in the file/input
	for input.Scan() {           // for each file
		counts[input.Text()]++ // increment the count if the words are found
	}
	// NOTE: ignoring potential errors from intput.Err()
}
