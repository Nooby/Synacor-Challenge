// Package vm implements the Execution and Memory Model of the Synacor VM.
package vm

import (
	"encoding/binary"
	"log"
	"os"
)

const (
	addressSpace     uint16 = 32768
	validNumberRange uint16 = 32776
)

// CPU represents the CPU that executes the instructions loaded into its memory.
type CPU struct {
	cursor     uint16
	consoleIn  chan uint16
	consoleOut chan uint16
	Memory     []uint16
	Registers  [8]uint16
	Stack      []uint16
}

// NewCPU constructs a new CPU.
func NewCPU(in chan uint16, out chan uint16) *CPU {
	c := CPU{consoleIn: in, consoleOut: out}
	return &c
}

// LoadImage loads a binary file into the main Memory of the CPU.
func (cpu *CPU) LoadImage(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	memorySize, err := imageSize(file)
	if err != nil {
		return err
	}

	cpu.Memory = make([]uint16, memorySize)

	var pi uint16
	for i := range cpu.Memory {
		err = binary.Read(file, binary.LittleEndian, &pi)
		if err != nil {
			return err
		}
		cpu.Memory[i] = pi
	}

	return nil
}

// Execute reads operations from the main Memory and interpretes them until the function
// encounters the Halt OP or a Fatal Error occurs.
func (cpu *CPU) Execute() {
	for {
		switch code := cpu.read(); code {
		case halt:
			cpu.halt()
		case set:
			cpu.set()
		case push:
			cpu.push()
		case pop:
			cpu.pop()
		case eq:
			cpu.eq()
		case gt:
			cpu.gt()
		case jmp:
			cpu.jmp()
		case jt:
			cpu.jt()
		case jf:
			cpu.jf()
		case add:
			cpu.add()
		case mult:
			cpu.mult()
		case mod:
			cpu.mod()
		case and:
			cpu.and()
		case or:
			cpu.or()
		case not:
			cpu.not()
		case rmem:
			cpu.rmem()
		case wmem:
			cpu.wmem()
		case call:
			cpu.call()
		case ret:
			cpu.ret()
		case out:
			cpu.out()
		case in:
			cpu.in()
		case noop: // Nothing to do.
		default:
			log.Fatal("Unknown OpCode")
		}
	}
}

func (cpu *CPU) read() uint16 {
	element := cpu.Memory[cpu.cursor]
	cpu.cursor++
	if element >= validNumberRange {
		log.Fatal("Value outside of Integer Range.")
	}
	return element
}

func (cpu *CPU) readAsValue() uint16 {
	element := cpu.read()
	if element >= addressSpace {
		register := element - addressSpace
		element = cpu.Registers[register]
	}
	return element
}

func (cpu *CPU) readAsRegister() uint16 {
	element := cpu.read()
	if element < addressSpace {
		log.Fatal(string(element) + " is not a Register.")
	}
	register := element - addressSpace
	return register
}

// imageSize calculates how many uint16 values need to be stored in VM Memory
func imageSize(file *os.File) (uint16, error) {
	fi, err := file.Stat()
	if err != nil {
		return 0, err
	}
	memorySize := uint16(fi.Size() / 2)
	return memorySize, nil
}
