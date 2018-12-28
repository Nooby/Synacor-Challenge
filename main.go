package main

import (
	"fmt"

	"github.com/Nooby/Synacor-Challenge/term"
	"github.com/Nooby/Synacor-Challenge/vm"
)

func main() {
	challengeBinaryPath := "challenge.bin"

	in, out := term.HandleConsole()

	cpu := vm.NewCPU(in, out)
	err := cpu.LoadImage(challengeBinaryPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	cpu.Execute()
}
