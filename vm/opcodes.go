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
func (c *CPU) halt() {
	os.Exit(0)
}

// set a b
// set register <a> to the value of <b>
func (c *CPU) set() {}

// push a
// push <a> onto the stack
func (c *CPU) push() {}

// pop a
// remove the top element from the stack and write it into <a>; empty stack = error
func (c *CPU) pop() {}

// eq a b c
// set <a> to 1 if <b> is equal to <c>; set it to 0 otherwise
func (c *CPU) eq() {}

// gt a b c
// set <a> to 1 if <b> is greater than <c>; set it to 0 otherwise
func (c *CPU) gt() {}

// jmp a
// jump to <a>
func (c *CPU) jmp() {
	value := c.read()
	c.cursor = value
}

// jt a b
// if <a> is nonzero, jump to <b>
func (c *CPU) jt() {}

// jf a b
// if <a> is zero, jump to <b>
func (c *CPU) jf() {}

// add a b c
// assign into <a> the sum of <b> and <c> (modulo 32768)
func (c *CPU) add() {}

// mult: 10 a b c
// store into <a> the product of <b> and <c> (modulo 32768)
func (c *CPU) mult() {}

// mod a b c
// store into <a> the remainder of <b> divided by <c>
func (c *CPU) mod() {}

// and a b c
// stores into <a> the bitwise and of <b> and <c>
func (c *CPU) and() {}

// or a b c
// stores into <a> the bitwise or of <b> and <c>
func (c *CPU) or() {}

// not a b
// stores 15-bit bitwise inverse of <b> in <a>
func (c *CPU) not() {}

// rmem a b
// read memory at address <b> and write it to <a>
func (c *CPU) rmem() {}

// wmem a b
// write the value from <b> into memory at address <a>
func (c *CPU) wmem() {}

// call a
// write the address of the next instruction to the stack and jump to <a>
func (c *CPU) call() {}

// ret: 18
// remove the top element from the stack and jump to it; empty stack = halt
func (c *CPU) ret() {}

// out a
// write the character represented by ascii code <a> to the terminal
func (c *CPU) out() {
	value := c.read()
	var text = string(value)
	fmt.Print(text)
}

// in a
// read a character from the terminal and write its ascii code to <a>
func (c *CPU) in() {}

