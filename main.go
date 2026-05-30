// Main package is required for executable Go programs
package main

import(
	// Buffered I/O functionality for reading effeciency
	"bufio"
	// Formatted I/O for printing output
	"fmt"
	// Basic I/O interfaces and utilities
	//"io"
	// OS functionality
	"os"
	"strings"
	"unicode/utf8"
	"flag"
)

// Takes a slice of strings of max width and prepends/appends
// margins on first and last lines, at start and end of each line,
// and retunrs a string w/ the contents of the balloon
func buildBalloon(lines []string, maxwidth int) string{
	// Stores the balloon border chars
	var borders []string

	// # of lines in the message
	count := len(lines)

	// Slice that will store all balloon lines
	var ret []string

	// Border chars used for diff positions in the balloon
	borders = []string{"/", "\\", "\\", "/", "|", "<", ">"}

	// Creates the top border using _s
	top := " " + strings.Repeat("_", maxwidth+2)

	// Creates the bottom border using -s
	bottom := " " + strings.Repeat("-", maxwidth+2)

	// Adds the top border to output
	ret = append(ret, top)

	// Special case
	if count == 1{
		// Creates single line balloon format: < text >
		s := fmt.Sprintf("%s %s %s", borders[5], lines[0], borders[6])

		// Adds the line to output
		ret = append(ret, s)
	// Apparently Go requires else to be on the same line as the }
	} else{
		// Creates the first line of multi line balloon: / text \
		s := fmt.Sprintf(`%s %s %s`, borders[0], lines[0], borders[1])

		// Adds the first line
		ret = append(ret, s)

		// Starts at the second line
		i := 1

		// Processes all middle lines
		for ; i < count - 1; i++{
			// Formats middle line: | text |
			s = fmt.Sprintf(`%s %s %s`, borders[4], lines[i], borders[4])
		
			// Adds the middle line
			ret = append(ret, s)
		}

		// Creates last line: \ text /
		s = fmt.Sprintf(`%s %s %s`, borders[2], lines[i], borders[3])

		// Adds last line
		ret = append(ret, s)
	}

	// Adds the bottom border
	ret = append(ret, bottom)

	// Joins all balloon lines together w/ newline chars
	return strings.Join(ret, "\n")
}

// Converts all tabs found in the strings
// found in the 'lines' slice to 4 spaces, to prevent
// misalignments in counting the runes
func tabsToSpaces(lines []string) []string{
	// Slice that will store modified strings
	var ret []string

	// Loops through every line
	for _, l := range lines{
		// Replaces every tab with 4 spaces
		l = strings.Replace(l, "\t", "    ", -1)

		// Stores the modified line
		ret = append(ret, l)
	}

	// Returns the updated slice
	return ret
}

// Given a slice of strings, returns the length of
// the string w/ max length
func calculateMaxWidth(lines []string) int{
	// Current max width found
	w := 0

	// Loops through all lines
	for _, l := range lines{
		// Counts Unicode chars in the line
		len := utf8.RuneCountInString(l)

		// Updates max width if this line is longer
		if len > w {
			w = len
		}
	}

	return w
}

// Takes a slice of strings and appends to each one a
// # of spaces needed to have the all the same # of runes
func normalizeStringsLength(lines []string, maxwidth int) []string{
	// Slice that will store the padded strings
	var ret []string

	// Loops through every line
	for _, l := range lines{
		// Adds spaces to make the line length = to maxwidth
		s := l + strings.Repeat(" ", maxwidth - utf8.RuneCountInString(l))

		// Stores the padded line
		ret = append(ret, s)
	}

	return ret
}

func printFigure(name string){
	// ASCII art cow
	var cow = `         \  ^__^
          \ (oo)\_______
	    (__)\       )\/\
	        ||----w |
	        ||     ||
		`

	var dinosaur = `         \                      .       .
          \                    / ` + "`" + `.   .' "
           \           .---.  <    > <    >  .---.
            \          |    \  \ - ~ ~ - /  /    |
          _____           ..-~             ~-..-~
         |     |   \~~~\\.'                    ` + "`" + `./~~~/
        ---------   \__/                         \__/
       .'  O    \     /               /       \  "
      (_____,    ` + "`" + `._.'               |         }  \/~~~/
       ` + "`" + `----.          /       }     |        /    \__/
             ` + "`" + `-.      |       /      |       /      ` + "`" + `. ,~~|
                 ~-.__|      /_ - ~ ^|      /- _      ` + "`" + `..-'
                      |     /        |     /     ~-.     ` + "`" + `-. _  _  _
                      |_____|        |_____|         ~ - . _ _ _ _ _>

    `

	switch name{
	case "cow":
		fmt.Println(cow)
	case "dinosaur":
		fmt.Println(dinosaur)
	default:
		fmt.Println("Unknown figure.")
	}
}

func main() {
	// Gets info about standard input
	// The 2nd return value is ignored using _
	info, _ := os.Stdin.Stat()

	// Checks if stdin is connected directly to a terminal instead of a pipe
	if info.Mode()&os.ModeCharDevice != 0{
		fmt.Println("The command is intended to work with pipes.")

		fmt.Println("Usage: fortune | gocowsay")

		// Exits early
		return
	}

	// Slice that will store all input lines
	var lines []string

	var figure string
	// Defined a command line flag "-f" that lets the user choose
	// which ASCII figure the user wants to choose
	flag.StringVar(&figure, "f", "cow", "the figure name. Valid values are `cow` and `dinosaur`")
	flag.Parse()

	// Creates a buffered reader that reads from standard input
	//reader := bufio.NewReader(os.Stdin)

	scanner := bufio.NewScanner(os.Stdin)

	// Starts an infinite loop that reads input 1 rune at a time
	for scanner.Scan(){
		// Reads a single rune from stdin
		// The 2nd return value is ignored
		//line, _, err := reader.ReadRune()

		// Checks if an error occurred & the error is EOF (no more input to read)
		//if err != nil && err == io.EOF{
			// Exits the loop when all input has been processed
			//break
		//}

		// Adds the character that was read to the output slice
		lines = append(lines, scanner.Text())
	}

	// Replaces tabs w/ spaces
	lines = tabsToSpaces(lines)

	// Finds the longest line length
	maxwidth := calculateMaxWidth(lines)

	// Pads all lines to the same width
	messages := normalizeStringsLength(lines, maxwidth)

	// Builds the speech balloon
	balloon := buildBalloon(messages, maxwidth)

	// Prints balloon
	fmt.Println(balloon)

	// Prints figure
	printFigure(figure)

	// Prints extra blank line
	fmt.Println()
}
