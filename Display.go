package main

import "fmt"

const (
	BUFFER_WIDTH  int = 64
	BUFFER_HEIGHT int = 32
	PIXEL_COUNT   int = BUFFER_WIDTH * BUFFER_HEIGHT
)

type Display struct {
	buffer [PIXEL_COUNT]uint32
}

func (display *Display) CLEAR_DISPLAY() {

	for i := 0; i < PIXEL_COUNT; i++ {
		display.buffer[i] = 0x00000000
	}
}

func (display *Display) SET_PIXEL(i int, j int, value uint32) {
	fmt.Print("  ", i*64+j, "  ", value, "\n")
	display.buffer[i*64+j] = value
}

func (display *Display) GET_PIXEL(i int, j int) uint32 {
	return display.buffer[i*64+j]
}
