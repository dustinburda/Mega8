package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	fmt.Println("Hello World!")

	opcode := uint16(0xF40A)
	var CHIP_8CPU CPU
	CHIP_8CPU.DECODE(opcode)

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
				if t.State == sdl.PRESSED {
					switch t.Keysym.Sym {
					case sdl.K_1:
						fmt.Println("1 was pressed!")
						keymap[0x0] = true
					case sdl.K_2:
						fmt.Println("2 was pressed!")
						keymap[0x1] = true
					case sdl.K_3:
						fmt.Println("3 was pressed!")
						keymap[0x2] = true
					case sdl.K_4:
						fmt.Println("4 was pressed!")
						keymap[0x3] = true
					case sdl.K_q:
						fmt.Println("Q was pressed!")
						keymap[0x4] = true
					case sdl.K_w:
						fmt.Println("W was pressed!")
						keymap[0x5] = true
					case sdl.K_e:
						fmt.Println("E was pressed!")
						keymap[0x6] = true
					case sdl.K_r:
						fmt.Println("R was pressed!")
						keymap[0x7] = true
					case sdl.K_a:
						fmt.Println("A was pressed!")
						keymap[0x8] = true
					case sdl.K_s:
						fmt.Println("S was pressed!")
						keymap[0x9] = true
					case sdl.K_d:
						fmt.Println("D was pressed!")
						keymap[0xA] = true
					case sdl.K_f:
						fmt.Println("F was pressed!")
						keymap[0xB] = true
					case sdl.K_z:
						fmt.Println("Z was pressed!")
						keymap[0xC] = true
					case sdl.K_x:
						fmt.Println("X was pressed!")
						keymap[0xD] = true
					case sdl.K_c:
						fmt.Println("C was pressed!")
						keymap[0xE] = true
					case sdl.K_v:
						fmt.Println("V was pressed!")
						keymap[0xF] = true
					}
				}
				if t.State == sdl.RELEASED {
					switch t.Keysym.Sym {
					case sdl.K_1:
						fmt.Println("1 was released!")
						keymap[0x0] = false
					case sdl.K_2:
						fmt.Println("2 was released!")
						keymap[0x1] = false
					case sdl.K_3:
						fmt.Println("3 was released!")
						keymap[0x2] = false
					case sdl.K_4:
						fmt.Println("4 was released!")
						keymap[0x3] = false
					case sdl.K_q:
						fmt.Println("Q was released!")
						keymap[0x4] = false
					case sdl.K_w:
						fmt.Println("W was released!")
						keymap[0x5] = false
					case sdl.K_e:
						fmt.Println("E was released!")
						keymap[0x6] = false
					case sdl.K_r:
						fmt.Println("R was released!")
						keymap[0x7] = false
					case sdl.K_a:
						fmt.Println("A was released!")
						keymap[0x8] = false
					case sdl.K_s:
						fmt.Println("S was released!")
						keymap[0x9] = false
					case sdl.K_d:
						fmt.Println("D was released!")
						keymap[0xA] = false
					case sdl.K_f:
						fmt.Println("F was released!")
						keymap[0xB] = false
					case sdl.K_z:
						fmt.Println("Z was released!")
						keymap[0xC] = false
					case sdl.K_x:
						fmt.Println("X was released!")
						keymap[0xD] = false
					case sdl.K_c:
						fmt.Println("C was released!")
						keymap[0xE] = false
					case sdl.K_v:
						fmt.Println("V was released!")
						keymap[0xF] = false
					}
				}
			}
		}
	}
}
