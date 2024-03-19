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

	rect := sdl.Rect{0, 0, 200, 200}
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
					case sdl.K_2:
						fmt.Println("2 was pressed!")
					case sdl.K_3:
						fmt.Println("3 was pressed!")
					case sdl.K_4:
						fmt.Println("4 was pressed!")
					case sdl.K_q:
						fmt.Println("Q was pressed!")
					case sdl.K_w:
						fmt.Println("W was pressed!")
					case sdl.K_e:
						fmt.Println("E was pressed!")
					case sdl.K_r:
						fmt.Println("R was pressed!")
					case sdl.K_a:
						fmt.Println("A was pressed!")
					case sdl.K_s:
						fmt.Println("S was pressed!")
					case sdl.K_d:
						fmt.Println("D was pressed!")
					case sdl.K_f:
						fmt.Println("F was pressed!")
					case sdl.K_z:
						fmt.Println("Z was pressed!")
					case sdl.K_x:
						fmt.Println("X was pressed!")
					case sdl.K_c:
						fmt.Println("C was pressed!")
					case sdl.K_v:
						fmt.Println("V was pressed!")
					}
				}
			}
		}
	}
}
