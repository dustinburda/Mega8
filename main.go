package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	fmt.Println("Hello World!")

	// opcode := uint16(0xF40A)
	// var CHIP_8CPU CPU
	// CHIP_8CPU.DECODE(opcode)

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	var display [64][32]byte
	for i := 0; i < 64; i++ {
		for j := 0; j < 32; j++ {
			if i+j%2 == 0 {
				display[i][j] = 1
			}
		}
	}

	rect := sdl.Rect{0, 0, 780, 580}
	colour := sdl.Color{R: 255, G: 0, B: 255, A: 255} // purple
	pixel := sdl.MapRGBA(surface.Format, colour.R, colour.G, colour.B, colour.A)
	surface.FillRect(&rect, pixel)
	window.UpdateSurface()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			case *sdl.KeyboardEvent:
				handle_input(t, &running)
			}
		}
	}
}

func handle_input(t *sdl.KeyboardEvent, running *bool) {
	if t.State == sdl.PRESSED {
		switch t.Keysym.Sym {
		case sdl.K_ESCAPE:
			fmt.Println("Escape was pressed!\nQuit")
			*running = false
		case sdl.K_x:
			fmt.Println("X[0x0] was pressed!")
			keymap[0x0] = true
		case sdl.K_1:
			fmt.Println("1[0x1] was pressed!")
			keymap[0x1] = true
		case sdl.K_2:
			fmt.Println("2[0x2] was pressed!")
			keymap[0x2] = true
		case sdl.K_3:
			fmt.Println("3[0x3] was pressed!")
			keymap[0x3] = true
		case sdl.K_4:
			fmt.Println("4[0xC] was pressed!")
			keymap[0xC] = true
		case sdl.K_q:
			fmt.Println("Q[0x4] was pressed!")
			keymap[0x4] = true
		case sdl.K_w:
			fmt.Println("W[0x5] was pressed!")
			keymap[0x5] = true
		case sdl.K_e:
			fmt.Println("E[0x6] was pressed!")
			keymap[0x6] = true
		case sdl.K_r:
			fmt.Println("R[0xD] was pressed!")
			keymap[0xD] = true
		case sdl.K_a:
			fmt.Println("A[0x7] was pressed!")
			keymap[0x7] = true
		case sdl.K_s:
			fmt.Println("S[0x8] was pressed!")
			keymap[0x8] = true
		case sdl.K_d:
			fmt.Println("D[0x9] was pressed!")
			keymap[0x9] = true
		case sdl.K_f:
			fmt.Println("F[0xE] was pressed!")
			keymap[0xE] = true
		case sdl.K_z:
			fmt.Println("Z[0xA] was pressed!")
			keymap[0xA] = true
		case sdl.K_c:
			fmt.Println("C[0xB] was pressed!")
			keymap[0xB] = true
		case sdl.K_v:
			fmt.Println("V[0xF] was pressed!")
			keymap[0xF] = true
		}
	}
	if t.State == sdl.RELEASED {
		switch t.Keysym.Sym {
		case sdl.K_x:
			fmt.Println("X[0x0] was released!")
			keymap[0x0] = false
		case sdl.K_1:
			fmt.Println("1[0x1] was released!")
			keymap[0x1] = false
		case sdl.K_2:
			fmt.Println("2[0x2] was released!")
			keymap[0x2] = false
		case sdl.K_3:
			fmt.Println("3[0x3] was released!")
			keymap[0x3] = false
		case sdl.K_4:
			fmt.Println("4[0xC] was released!")
			keymap[0xC] = false
		case sdl.K_q:
			fmt.Println("Q[0x4] was released!")
			keymap[0x4] = false
		case sdl.K_w:
			fmt.Println("W[0x5] was released!")
			keymap[0x5] = false
		case sdl.K_e:
			fmt.Println("E[0x6] was released!")
			keymap[0x6] = false
		case sdl.K_r:
			fmt.Println("R[0xD] was released!")
			keymap[0xD] = false
		case sdl.K_a:
			fmt.Println("A[0x7] was released!")
			keymap[0x7] = false
		case sdl.K_s:
			fmt.Println("S[0x8] was released!")
			keymap[0x8] = false
		case sdl.K_d:
			fmt.Println("D[0x9] was released!")
			keymap[0x9] = false
		case sdl.K_f:
			fmt.Println("F[0xE] was released!")
			keymap[0xE] = false
		case sdl.K_z:
			fmt.Println("Z[0xA] was released!")
			keymap[0xA] = false
		case sdl.K_c:
			fmt.Println("C[0xB] was released!")
			keymap[0xB] = false
		case sdl.K_v:
			fmt.Println("V[0xF] was released!")
			keymap[0xF] = false
		}
	}
}
