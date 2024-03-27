package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var keymap = make(map[int]bool)

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
