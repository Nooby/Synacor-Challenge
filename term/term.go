package term

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// HandleBatchInput reads the content of the file at 'path' to the chan 'in'.
// After EOF of the file HandleBatchInput reads Stdin to the chan.
// Lines starting with '#' and empty lines will be ignored.
func HandleBatchInput(in chan uint16, path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("Unable to open Batch Input: %v", err)
	}
	readTo(in, f)
	readTo(in, os.Stdin)
}

// HandleConsoleInput reads Stdin to the chanel 'in'.
// Lines starting with '#' and empty lines will be ignored.
func HandleConsoleInput(in chan uint16) {
	readTo(in, os.Stdin)
}

// HandleOutput writes chan 'out' to Stdout.
func HandleOutput(out chan uint16) {
	var text string
	for c := range out {
		text = string(c)
		fmt.Print(text)
	}
}

func readTo(in chan uint16, reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) <= 0 || text[0] == '#' {
			continue
		}
		for ic := 0; ic < len(text); ic++ {
			value := text[ic]
			if value >= 128 {
				log.Fatal("Only ASCII Runes Supported.")
			}
			in <- uint16(value)
		}
		in <- 10 // Scanner removes the Line Feed; Re Adding it.
	}
}
