package main

import "fmt"

const (
	BUFFER_WIDTH  int = 64
	BUFFER_HEIGHT int = 32
)

type Display struct {
	buffer [BUFFER_WIDTH * BUFFER_HEIGHT]uint32
}

func (display *Display) CLEAR_DISPLAY() {

	for i := 0; i < BUFFER_WIDTH*BUFFER_HEIGHT; i++ {
		display.buffer[i] = 0x00000000
	}
}

func (display *Display) SET_PIXEL(i int, j int, value uint32) {
	fmt.Print("  ", i*64+j, "  ", value, "\n")
	display.buffer[i*64+j] = value
}
