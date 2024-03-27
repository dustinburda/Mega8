package main

import (
	"fmt"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	WIDTH  int32 = 1200
	HEIGHT int32 = 600
)

func main() {
	opcode := uint16(0xF40A)
	var CHIP_8CPU CPU
	CHIP_8CPU.DECODE(opcode)

	var d Display
	d.CLEAR_DISPLAY()
	CHIP_8CPU.display = &d

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("CHIP-8 Emulator",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		WIDTH, HEIGHT, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, 0)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	// var display [64 * 32]uint32
	for i := 0; i < 64*32; i++ {
		x := i / 64
		y := i % 64

		fmt.Print(uint8(x), "  ", uint8(y))
		if (x+y)%2 == 1 {
			CHIP_8CPU.display.SET_PIXEL(x, y, 0x00FFFFFF)
		} else {
			fmt.Println()
		}
	}

	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, 64, 32)
	if err != nil {
		panic(err)
	}
	defer texture.Destroy()

	// reader := bufio.NewReader(os.Stdin)

	running := true
	for running {
		// fmt.Println("Select which Pixel (0 - 2048) you would like to light up: ")

		// PixelNumStr, _ := reader.ReadString('\n')
		// PixelNumStr = strings.Split(PixelNumStr, "\n")[0]

		// PixelNum, err := strconv.Atoi(PixelNumStr)
		// if err != nil {
		// 	panic(err)
		// }

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

		// CHIP_8CPU.display.buffer[PixelNum] = 0x00FFFFFF

		texture.Update(nil, unsafe.Pointer(&CHIP_8CPU.display.buffer[0]), 64*int(unsafe.Sizeof(uint32(0))))
		renderer.Clear()
		renderer.Copy(texture, nil, nil)
		renderer.Present()
	}
}
