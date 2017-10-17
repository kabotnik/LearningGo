// Echo3 prints its command-line arguments
package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	args := strings.Join(os.Args[1:], " ")
	fmt.Println(args)

	totalTime := time.Since(start).Seconds()
	fmt.Println("Using string.Join()\t: %d sec", totalTime)

	start = time.Now()
	for key, arg := range os.Args[1:] {
		fmt.Println(strconv.Itoa(key) + " : " + arg)
	}
	totalTime = time.Since(start).Seconds()
	fmt.Println("Using iterator\t: %d sec", totalTime)
}