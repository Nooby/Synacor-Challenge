package vm

import (
	"os"
	"fmt"
	"encoding/binary"
)

const AdressSpace uint16 = 32768

type CPU struct {
	cursor uint16
	Memory []uint16
	Registers [8]uint16
	Stack []uint16
}

func (c *CPU) LoadImage(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	memorySize, err := imageSize(file)
	if err != nil {
		return err
	}

	c.Memory = make([]uint16, memorySize)

	var pi uint16
	for i := range c.Memory {
		err = binary.Read(file, binary.LittleEndian, &pi)
		if err != nil {
			return err
		}
		c.Memory[i] = pi
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

func (c *CPU) Execute() {
	for {
		switch code := c.read(); code {
		case halt: c.halt()
		case jmp: c.jmp()
		case out: c.out()
		case noop: 
		default:
			fmt.Printf("OpCode Not Implemented: %v", code)

		}
	}
}

func (c *CPU) read() uint16 {
	value := c.Memory[c.cursor]
	c.cursor += 1
	return value
}

