package vm

import (
	"os"
	"fmt"
	"encoding/binary"
)

const (
	addressSpace uint16 = 32768
	validNumberRange uint16 = 32776
)

type CPU struct {
	cursor uint16
	Memory []uint16
	Registers [8]uint16
	Stack []uint16
}

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

// imageSize calculates how many uint16 values need to be stored in VM Memory
func imageSize(file *os.File) (uint16, error) {
	fi, err := file.Stat()
	if err != nil {
		return 0, err
	}
	memorySize := uint16(fi.Size() / 2)
	return memorySize, nil
}

func (cpu *CPU) Execute() {
	for {
		switch code := cpu.read(); code {
		case halt: cpu.halt()
		case set: cpu.set()
		case push: cpu.push()
		case pop: cpu.pop()
		case eq: cpu.eq()
		case gt: cpu.gt()
		case jmp: cpu.jmp()
		case jt: cpu.jt()
		case jf: cpu.jf()
		case add: cpu.add()
		case mult: cpu.mult()
		case mod: cpu.mod()
		case and: cpu.and()
		case or: cpu.or()
		case not: cpu.not()
		case rmem: cpu.rmem()
		case wmem: cpu.wmem()
		case call: cpu.call()
		case ret: cpu.ret()
		case out: cpu.out()
		case noop: 
		default:
			fmt.Printf("OpCode Not Implemented: %v", code)
			return
		}
	}
}

func (cpu *CPU) read() uint16 {
	element := cpu.Memory[cpu.cursor]
	cpu.cursor += 1
	if element >= validNumberRange {
		panic("Value outside of Integer Range.")
		
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
	if element >= addressSpace {
		register := element - addressSpace
		return register
	}
	panic(string(element) + " is not a Register.")
}
