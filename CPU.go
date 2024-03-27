package main

import (
	// "bufio"
	"errors"
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

func (cpu *CPU) LOAD_ROM(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, errors.New("failed to load ROM")
	}
	defer file.Close()

	bytes_read, err := file.Read(cpu.memory[PROGRAM_START:])
	if err != nil {
		return 0, errors.New("failed to read ROM")
	}

	return bytes_read, nil

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

	// TODO: learn more about timers
	cpu.delay_timer = 60
	cpu.sound_timer = 60
}

// TODO: initialize display here
// TODO: initialize
func (cpu *CPU) CPU_INIT() {
	fmt.Println("INIT")
}

// Single Fetch-Decode-Execute Cycle
func (cpu *CPU) CYCLE() {
	opcode := cpu.FETCH() // OPCODE
	cpu.DECODE(opcode)    // DECODE + EXECUTE

	// TODO: add timer updates
}

func (cpu *CPU) MAIN_LOOP() {
	for {
		cpu.CYCLE()
	}
}

func (cpu *CPU) FETCH() uint16 {
	var op uint16 = 0x0

	// stored big endian
	word_first_half := cpu.memory[cpu.PC]
	word_second_half := cpu.memory[cpu.PC+1]

	op |= uint16(word_first_half) << 8
	op |= uint16(word_second_half)

	cpu.PC += 2

	return op
}

func (cpu *CPU) DECODE(opcode uint16) func(uint16) {
	nibble_one := (opcode & 0xF000) >> 12
	nibble_two := (opcode & 0x0F00) >> 8
	nibble_three := (opcode & 0x00F0) >> 4
	nibble_four := opcode & 0x000F

	fmt.Println(nibble_one, nibble_two, nibble_three, nibble_four)

	switch nibble_one {
	case 0x0:
		switch nibble_four {
		case 0x0:
			cpu.EXECUTE_0x00E0()
		case 0xE:
			cpu.EXECUTE_0x00EE()
		}
	case 0x1:
		cpu.EXECUTE_0x1NNN(opcode)
	case 0x2:
		cpu.EXECUTE_0x2NNN()
	case 0x3:
		cpu.EXECUTE_0x3XNN()
	case 0x4:
		cpu.EXECUTE_0x4XNN()
	case 0x5:
		cpu.EXECUTE_0x5XY0()
	case 0x6:
		cpu.EXECUTE_0x6XNN(opcode)
	case 0x7:
		cpu.EXECUTE_0x7XNN(opcode)
	case 0x8:
		switch nibble_four {
		case 0x0:
			cpu.EXECUTE_0x8XY0()
		case 0x1:
			cpu.EXECUTE_0x8XY1()
		case 0x2:
			cpu.EXECUTE_0x8XY2()
		case 0x3:
			cpu.EXECUTE_0x8XY3()
		case 0x4:
			cpu.EXECUTE_0x8XY4()
		case 0x5:
			cpu.EXECUTE_0x8XY5()
		case 0x6:
			cpu.EXECUTE_0x8XY6()
		case 0x7:
			cpu.EXECUTE_0x8XY7()
		case 0xE:
			cpu.EXECUTE_0x8XYE()
		}
	case 0x9:
		cpu.EXECUTE_0x9XY0()
	case 0xA:
		cpu.EXECUTE_0xANNN(opcode)
	case 0xB:
		cpu.EXECUTE_0xBNNN()
	case 0xC:
		cpu.EXECUTE_0xCXNN()
	case 0xD:
		cpu.EXECUTE_0xDXYN(opcode)
	case 0xE:
		switch nibble_three<<4 | nibble_four {
		case 0x9E:
			cpu.EXECUTE_0xEX9E()
		case 0xA1:
			cpu.EXECUTE_0xEXA1()
		}
	case 0xF:
		switch nibble_three<<4 | nibble_four {
		case 0x07:
			cpu.EXECUTE_0xFX07()
		case 0x0A:
			cpu.EXECUTE_0xFX0A()
		case 0x15:
			cpu.EXECUTE_0xFX15()
		case 0x18:
			cpu.EXECUTE_0xFX18()
		case 0x1E:
			cpu.EXECUTE_0xFX1E()
		case 0x29:
			cpu.EXECUTE_0xFX29()
		case 0x33:
			cpu.EXECUTE_0xFX33()
		case 0x55:
			cpu.EXECUTE_0xFX55()
		case 0x65:
			cpu.EXECUTE_0xFX65()
		}
	default:
		fmt.Println("OPCODE, ", opcode, " not recognized")
	}

	return nil
}
