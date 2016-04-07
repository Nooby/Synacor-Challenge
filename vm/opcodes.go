package vm

import (
	"os"
	"fmt"
)

const (
	halt uint16 = iota
	set
	push
	pop
	eq
	gt
	jmp
	jt
	jf
	add
	mult
	mod
	and
	or
	not
	rmem
	wmem
	call
	ret
	out
	in
	noop
)

// halt stops execution and terminates the program
func (cpu *CPU) halt() {
	os.Exit(0)
}

// set a b
// set register <a> to the value of <b>
func (cpu *CPU) set() {
	a := cpu.readAsRegister()
	b := cpu.readAsValue()
	cpu.Registers[a] = b
}

// push a
// push <a> onto the stack
func (cpu *CPU) push() {
	a := cpu.readAsValue()
	cpu.Stack = append(cpu.Stack, a)
}

// pop a
// remove the top element from the stack and write it into <a>; empty stack = error
func (cpu *CPU) pop() {
	a := cpu.readAsRegister()
	element := cpu.Stack[len(cpu.Stack) - 1]
	cpu.Stack = cpu.Stack[:len(cpu.Stack) - 1]
	cpu.Registers[a] = element
}

// eq a b c
// set <a> to 1 if <b> is equal to <c>; set it to 0 otherwise
func (cpu *CPU) eq() {
	a := cpu.readAsRegister()
	b := cpu.readAsValue()
	c := cpu.readAsValue()

	if b == c {
		cpu.Registers[a] = 1
	} else {
		cpu.Registers[a] = 0
	}
}

// gt a b c
// set <a> to 1 if <b> is greater than <c>; set it to 0 otherwise
func (cpu *CPU) gt() {
	a := cpu.readAsRegister()
	b := cpu.readAsValue()
	c := cpu.readAsValue()

	if b > c {
		cpu.Registers[a] = 1
	} else {
		cpu.Registers[a] = 0
	}
}

// jmp a
// jump to <a>
func (cpu *CPU) jmp() {
	a := cpu.readAsValue()
	cpu.cursor = a
}

// jt a b
// if <a> is nonzero, jump to <b>
func (cpu *CPU) jt() {
	a := cpu.readAsValue()
	b := cpu.readAsValue()
	if a != 0 {
		cpu.cursor = b
	}
}

// jf a b
// if <a> is zero, jump to <b>
func (cpu *CPU) jf() {
	a := cpu.readAsValue()
	b := cpu.readAsValue()
	if a == 0 {
		cpu.cursor = b
	}
}

// add a b c
// assign into <a> the sum of <b> and <c> (modulo 32768)
func (cpu *CPU) add() {
	a := cpu.readAsRegister()
	b := cpu.readAsValue()
	c := cpu.readAsValue()

	sum := b + c
	if sum >= addressSpace {
		sum = sum % addressSpace
	}

	cpu.Registers[a] = sum
}

// mult: 10 a b c
// store into <a> the product of <b> and <c> (modulo 32768)
func (cpu *CPU) mult() {
	a := cpu.readAsRegister()
	b := cpu.readAsValue()
	c := cpu.readAsValue()

	prod := b * c
	if prod >= addressSpace {
		prod = prod % addressSpace
	}

	cpu.Registers[a] = prod
}

// mod a b c
// store into <a> the remainder of <b> divided by <c>
func (cpu *CPU) mod() {
	a := cpu.readAsRegister()
	b := cpu.readAsValue()
	c := cpu.readAsValue()

	rem := b % c
	if rem >= addressSpace {
		rem = rem % addressSpace
	}
	cpu.Registers[a] = rem
}

// and a b c
// stores into <a> the bitwise and of <b> and <c>
func (cpu *CPU) and() {
	a := cpu.readAsRegister()
	b := cpu.readAsValue()
	c := cpu.readAsValue()

	res := b & c
	if res >= addressSpace {
		res = res % addressSpace
	}
	cpu.Registers[a] = res
}

// or a b c
// stores into <a> the bitwise or of <b> and <c>
func (cpu *CPU) or() {
	a := cpu.readAsRegister()
	b := cpu.readAsValue()
	c := cpu.readAsValue()

	res := b | c
	if res >= addressSpace {
		res = res % addressSpace
	}
	cpu.Registers[a] = res
}

// not a b
// stores 15-bit bitwise inverse of <b> in <a>
func (cpu *CPU) not() {
	a := cpu.readAsRegister()
	b := cpu.readAsValue()
	res := ^b
	if res >= addressSpace {
		res = res % addressSpace
	}
	cpu.Registers[a] = res
}

// rmem a b
// read memory at address <b> and write it to <a>
func (cpu *CPU) rmem() {
	a := cpu.readAsRegister()
	b := cpu.readAsValue()
	cpu.Registers[a] = cpu.Memory[b]
}

// wmem a b
// write the value from <b> into memory at address <a>
func (cpu *CPU) wmem() {
	a := cpu.readAsValue()
	b := cpu.readAsValue()
	cpu.Memory[a] = b
}

// call a
// write the address of the next instruction to the stack and jump to <a>
func (cpu *CPU) call() {
	a := cpu.readAsValue()
	cpu.Stack = append(cpu.Stack, cpu.cursor)
	cpu.cursor = a
}

// ret
// remove the top element from the stack and jump to it; empty stack = halt
func (cpu *CPU) ret() {
	element := cpu.Stack[len(cpu.Stack) - 1]
	cpu.Stack = cpu.Stack[:len(cpu.Stack) - 1]
	cpu.cursor = element
}

// out a
// write the character represented by ascii code <a> to the terminal
func (cpu *CPU) out() {
	a:= cpu.readAsValue()
	var text = string(a)
	fmt.Print(text)
}

// in a
// read a character from the terminal and write its ascii code to <a>
func (cpu *CPU) in() {}

