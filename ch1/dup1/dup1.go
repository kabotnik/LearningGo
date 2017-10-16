// Prints the text of each line that appears more than
// once in the standard input, preceded by its count
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
		// above is equivalent to the following:
		// line := input.Text()
		// counts[line] = counts[line] + 1
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}