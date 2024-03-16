package main

import (
	// "bufio"
	"fmt"
	"os"
	// "io"
)

const PROGRAM_START uint16 = 0x200

type CPU struct {
	memory         [0xFFF]uint8
	data_registers [16]uint8
	register_I     uint16
	PC             uint16
	stack          []uint16

	delay_timer uint8
	sound_timer uint8

	display *Display
}

func (cpu *CPU) LOAD_ROM(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Failed to load ROM!")
		return
	}
	defer file.Close()

	file.Read(cpu.memory[PROGRAM_START:])

}

func (cpu *CPU) CPU_RESET() {
	for i := 0; i < 0xFFF; i++ {
		cpu.memory[i] = 0
	}

	for i := 0; i < 16; i++ {
		cpu.data_registers[i] = 0
	}

	cpu.register_I = 0
	cpu.PC = PROGRAM_START
	clear(cpu.stack) //  = nil   Setting slice to nil will release the underlying memory to the garbage collection

	cpu.display.CLEAR_DISPLAY()
}

func (cpu *CPU) CPU_INIT() {
	fmt.Println("INIT")
}

func (cpu *CPU) MAIN_LOOP() {
	for {
		fmt.Println("RUNNING")
	}
}
