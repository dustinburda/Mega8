package main

import (
	// "bufio"
	"fmt"
	// "io"
	// "os"
)

type CPU struct {
	memory         [0xFFF]uint8
	data_registers [16]uint8
	register_I     uint16
	PC             uint16
	stack          []uint16
}

func (cpu *CPU) LOAD_ROM() {
	fmt.Println("LOADING ROM")
}

func (cpu *CPU) CPU_INIT() {
	fmt.Println("INIT")
}
