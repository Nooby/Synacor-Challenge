package main

import (
	"fmt"
	"github.com/nooby/synacor-challenge/vm"
)

func main() {
	challengeBinaryPath := "challenge.bin"

	var cpu = vm.CPU{}
	err := cpu.LoadImage(challengeBinaryPath)
	if err != err {
		fmt.Println(err)
		return
	}
	cpu.Execute()
}
