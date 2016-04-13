package term

import (
	"os"
	"fmt"
	"bufio"
)

func HandleConsole() (consoleInput, consoleOutput chan uint16) {
	consoleInput = make(chan uint16)
	consoleOutput = make(chan uint16)

	go HandleIn(consoleInput)
	go HandleOut(consoleOutput)

	return consoleInput, consoleOutput
}

func HandleIn(in chan uint16) {
	var text string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text = scanner.Text()
		for ic := 0; ic < len(text); ic++ {
			value := text[ic]
			if value >= 128 {
				panic("Only ASCII Runes Supported.")
			}
			in<- uint16(value)
		}
		in<- 10	// Scanner removes the Line Feed; Re Adding it.
	}
}

func HandleOut(out chan uint16) {
	var text string
	for c := range out {
		text = string(c)
		fmt.Print(text)
	}
}
