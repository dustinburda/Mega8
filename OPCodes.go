package main

import "fmt"

// @brief: Clears display
func (cpu *CPU) EXECUTE_0x00E0() {
	fmt.Println("Executing 0x00E0")
	cpu.display.CLEAR_DISPLAY()
}

func (cpu *CPU) EXECUTE_0x00EE() {
	fmt.Println("Executing 0x00E0")
}

func (cpu *CPU) EXECUTE_0x1NNN(opcode uint16) {
	fmt.Println("Executing 0x1NNN")
	cpu.PC = (opcode & 0x0FFF)

}

func (cpu *CPU) EXECUTE_0x2NNN() {
	fmt.Println("Executing 0x2NNN")
}

func (cpu *CPU) EXECUTE_0x3XNN() {
	fmt.Println("Executing 0x3XNN")
}

func (cpu *CPU) EXECUTE_0x4XNN() {
	fmt.Println("Executing 0x4XNN")
}

func (cpu *CPU) EXECUTE_0x5XY0() {
	fmt.Println("Executing 0x5XY0")
}

// @brief: sets register Vx to value NN
// @param: opcode to be executed
func (cpu *CPU) EXECUTE_0x6XNN(opcode uint16) {
	fmt.Println("Executing 0x6XNN")

	Vx := (opcode & 0x0F00) >> 8
	NN := (opcode & 0x00FF)

	cpu.data_registers[Vx] = uint8(NN)
}

// @brief: adds NN to the register Vx
// @param: opcode to be executed
func (cpu *CPU) EXECUTE_0x7XNN(opcode uint16) {
	fmt.Println("Executing 0x7XNN")

	Vx := (opcode & 0x0F00) >> 8
	NN := (opcode & 0x00FF)

	cpu.data_registers[Vx] += uint8(NN)

}

func (cpu *CPU) EXECUTE_0x8XY0() {
	fmt.Println("Executing 0x8XY0")
}
func (cpu *CPU) EXECUTE_0x8XY1() {
	fmt.Println("Executing 0x8XY1")
}

func (cpu *CPU) EXECUTE_0x8XY2() {
	fmt.Println("Executing 0x8XY2")
}

func (cpu *CPU) EXECUTE_0x8XY3() {
	fmt.Println("Executing 0x8XY3")
}

func (cpu *CPU) EXECUTE_0x8XY4() {
	fmt.Println("Executing 0x8XY4")
}

func (cpu *CPU) EXECUTE_0x8XY5() {
	fmt.Println("Executing 0x8XY5")
}

func (cpu *CPU) EXECUTE_0x8XY6() {
	fmt.Println("Executing 0x8XY6")
}

func (cpu *CPU) EXECUTE_0x8XY7() {
	fmt.Println("Executing 0x8XY7")
}

func (cpu *CPU) EXECUTE_0x8XYE() {
	fmt.Println("Executing 0x8XYE")
}

func (cpu *CPU) EXECUTE_0x9XY0() {
	fmt.Println("Executing 0x9XY0")
}

func (cpu *CPU) EXECUTE_0xANNN(opcode uint16) {
	fmt.Println("Executing 0xANNN")

	NNN := (opcode & 0x0FFF)

	cpu.register_I = NNN
}

func (cpu *CPU) EXECUTE_0xBNNN() {
	fmt.Println("Executing 0xBNNN")
}

func (cpu *CPU) EXECUTE_0xCXNN() {
	fmt.Println("Executing 0xCXNN")
}

// @brief: draw a sprite of height N at X,Y
// @param:
func (cpu *CPU) EXECUTE_0xDXYN(opcode uint16) {
	fmt.Println("Executing 0xDXYN")

	Vx := (opcode & 0x0F00) >> 8
	Vy := (opcode & 0x00F0) >> 4
	N := (opcode & 0x000F)

	x := cpu.data_registers[Vx] % 64
	y := cpu.data_registers[Vy] % 32

	cpu.data_registers[0xF] = 0

	for i := 0; i < int(N); i++ {
		_ = cpu.memory[cpu.register_I+uint16(i)]

		for j := 0; j < 8; j++ {

			x += 1

			// TODO:

			if x >= 64 {
				break
			}
		}
		y += 1

		if y >= 32 {
			break
		}

	}
}

func (cpu *CPU) EXECUTE_0xEX9E() {
	fmt.Println("Executing 0xEX9E")
}

func (cpu *CPU) EXECUTE_0xEXA1() {
	fmt.Println("Executing 0xEXA1")
}

func (cpu *CPU) EXECUTE_0xFX07() {
	fmt.Println("Executing 0xFX07")
}

func (cpu *CPU) EXECUTE_0xFX0A() {
	fmt.Println("Executing 0xFX0A")
}

func (cpu *CPU) EXECUTE_0xFX15() {
	fmt.Println("Executing 0xFX15")
}

func (cpu *CPU) EXECUTE_0xFX18() {
	fmt.Println("Executing 0xFX18")
}

func (cpu *CPU) EXECUTE_0xFX1E() {
	fmt.Println("Executing 0xFX1E")
}

func (cpu *CPU) EXECUTE_0xFX29() {
	fmt.Println("Executing 0xFX29")
}

func (cpu *CPU) EXECUTE_0xFX33() {
	fmt.Println("Executing 0xFX33")
}

func (cpu *CPU) EXECUTE_0xFX55() {
	fmt.Println("Executing 0xFX55")
}

func (cpu *CPU) EXECUTE_0xFX65() {
	fmt.Println("Executing 0xFX65")
}
