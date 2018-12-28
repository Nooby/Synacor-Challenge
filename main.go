package main

import (
	"flag"
	"fmt"

	"github.com/Nooby/Synacor-Challenge/term"
	"github.com/Nooby/Synacor-Challenge/vm"
)

const (
	defaultChallengePath = "challenge.bin"
	challengePathUsage   = "path to challenge binary"
	commandInputUsage    = "path to batch commands file"
)

var challengeBin string
var commandInput string

func init() {
	flag.StringVar(&challengeBin, "bin",
		defaultChallengePath, challengePathUsage)
	flag.StringVar(&challengeBin, "b",
		defaultChallengePath, challengePathUsage+" (shorthand)")
	flag.StringVar(&commandInput, "input", "", commandInputUsage)
	flag.StringVar(&commandInput, "i", "", commandInputUsage+" (shorthand)")
}

func main() {
	flag.Parse()

	in := make(chan uint16)
	out := make(chan uint16)

	if len(commandInput) > 0 {
		go term.HandleBatchInput(in, commandInput)
	} else {
		go term.HandleConsoleInput(in)
	}

	go term.HandleOutput(out)

	cpu := vm.NewCPU(in, out)
	err := cpu.LoadImage(challengeBin)
	if err != nil {
		fmt.Println(err)
		return
	}

	cpu.Execute()
}
