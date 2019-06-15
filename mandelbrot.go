package main
import (
	"math"
)
import "github.com/veandco/go-sdl2/sdl"

func map_values(value int, in_min, in_max, out_min, out_max float64) float64 {
	return (float64(value) - in_min) * (out_max - out_min) / (in_max - in_min) + out_min
}

func main() {
	const WIDTH = 800
	const HEIGTH = 800
	const MAX_ITERATIONS = 200
	
	var min float64 = -2.0
	var max float64 = 2.0
	
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	// var event *sdl.Event

	// Cria window and renderer
	window, renderer, err := sdl.CreateWindowAndRenderer(1400, 900, 0);
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	renderer.SetLogicalSize(WIDTH, HEIGTH)

	for {
		// O PollEvent é necessário para avisar o OS e não deixar a janela irresponsiva
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				break
			}
		}
		renderer.Clear()

		// Algoritmo de criação do Mandelbrot
		for x := 0; x < WIDTH; x++ {
			for y := 0; y < HEIGTH; y++ {
				a := map_values(x, 0, WIDTH, min, max)
				b := map_values(y, 0, HEIGTH, min, max)

				var ai = a
				var bi = b

				n := 0

				for i := 0; i < MAX_ITERATIONS; i++ {
					var a1 float64
					var b1 float64

					a1 = a*a - b*b
					b1 = 2*a*b

					a = a1 + ai
					b = b1 + bi

					if (a+b) > 2 {
						break
					}

					n = n + 1
				}

				brigth := map_values(n, 0, MAX_ITERATIONS, 0, 255)

				// Se a função não divergiu, então o ponto pertence ao conjunto de Mandelbrot
				if n == MAX_ITERATIONS || brigth <= 20 {
					brigth = 0
				}

				// Cores
				red := map_values(int(brigth*brigth), 0, 255*255, 0, 255)
				green := brigth
				blue := map_values(int(math.Sqrt(brigth)), 0, math.Sqrt(255), 0, 255)

				renderer.SetDrawColor(uint8(red), uint8(green), uint8(blue), 255)
				renderer.DrawPoint(int32(x), int32(y))
			}
		}

		renderer.Present()
	}
}