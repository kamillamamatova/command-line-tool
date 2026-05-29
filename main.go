// Main package is required for executable Go programs
package main

import (
	// Buffered I/O functionality for reading effeciency
	"bufio"
	// Formatted I/O for printing output
	"fmt"
	// Basic I/O interfaces and utilities
	"io"
	// OS functionality
	"os"
)
func main() {
	// Gets info about standard input
	// The 2nd return value is ignored using _
	info, _ := os.Stdin.Stat()

	// Checks if stdin is connected directly to a terminal instead of a pipe
	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")

		fmt.Println("Usage: fortune | gocowsay")

		// Exits early
		return
	}

	// Creates a buffered reader that reads from standard input
	reader := bufio.NewReader(os.Stdin)

	// Creates a slice of runes to store all chars read from stdin
	var output []rune

	// Starts an infinite loop that reads input 1 rune at a time
	for {
		// Reads a single rune from stdin
		// The 2nd return value is ignored
		input, _, err := reader.ReadRune()

		// Checks if an error occurred & the error is EOF (no more input to read)
		if err != nil && err == io.EOF {
			// Exits the loop when all input has been processed
			break
		}

		// Adds the character that was read to the output slice
		output = append(output, input)
	}

	// Loops through every rune stored in the output slice
	for j := 0; j < len(output); j++ {
		fmt.Printf("%c", output[j])
	}
}
