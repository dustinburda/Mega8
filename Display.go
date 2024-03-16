package main

type Display struct {
	buffer [64][32]byte
}

func (display *Display) CLEAR_DISPLAY() {
	for i := 0; i < 64; i++ {
		for j := 0; j < 32; j++ {
			display.buffer[i][j] = 0
		}
	}
}

func (display *Display) SET_PIXEL(i uint8, j uint8, value byte) {
	display.buffer[i][j] = value
}
