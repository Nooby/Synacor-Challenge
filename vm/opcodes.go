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

func (c *CPU) halt() {
	os.Exit(0)
}

func (c *CPU) set() {}
func (c *CPU) push() {}
func (c *CPU) pop() {}
func (c *CPU) eq() {}
func (c *CPU) gt() {}

func (c *CPU) jmp() {
	value := c.read()
	c.cursor = value
}

func (c *CPU) jt() {}
func (c *CPU) jf() {}
func (c *CPU) add() {}
func (c *CPU) mult() {}
func (c *CPU) mod() {}
func (c *CPU) and() {}
func (c *CPU) or() {}
func (c *CPU) not() {}
func (c *CPU) rmem() {}
func (c *CPU) wmem() {}
func (c *CPU) call() {}
func (c *CPU) ret() {}

func (c *CPU) out() {
	value := c.read()
	var text = string(value)
	fmt.Print(text)
}

func (c *CPU) in() {}

