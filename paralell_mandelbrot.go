package main
import (
	"math"
	"strconv"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)
import "github.com/veandco/go-sdl2/sdl"

func map_values(value int, in_min, in_max, out_min, out_max float64) float64 {
	return (float64(value) - in_min) * (out_max - out_min) / (in_max - in_min) + out_min
}

func printFractal(WIDTH, HEIGTH, maxIterations, count int, min, max float64, renderer *sdl.Renderer) {
	bounds := image.Rect(0, 0, WIDTH, HEIGTH)
	b_set := image.NewNRGBA(bounds)
	draw.Draw(b_set, bounds, image.NewUniform(color.Black), image.ZP, draw.Src)

	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGTH; y++ {
			a := map_values(x, 0, float64(WIDTH), min, max)
			b := map_values(y, 0, float64(HEIGTH), min, max)

			var ai = a
			var bi = b

			n := 0

			for i := 0; i < maxIterations; i++ {
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

			brigth := map_values(n, 0, float64(maxIterations), 0, 255)

			// Se a função não divergiu, então o ponto pertence ao conjunto de Mandelbrot
			if n == maxIterations || brigth <= 20 {
				brigth = 0
			}

			// Cores
			red := map_values(int(brigth*brigth), 0, 255*255, 0, 255)
			green := brigth
			blue := map_values(int(math.Sqrt(brigth)), 0, math.Sqrt(255), 0, 255)

			renderer.SetDrawColor(uint8(red), uint8(green), uint8(blue), 255)
			renderer.DrawPoint(int32(x), int32(y))

			b_set.Set(x, y, color.NRGBA{uint8(red),
				uint8(green), uint8(blue), 255})
		}
	}

	file_name := strconv.Itoa(count) + ".png"
	f, err := os.Create("png/" + file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = png.Encode(f, b_set); err != nil {
		fmt.Println(err)
	}
	if err = f.Close(); err != nil {
		fmt.Println(err)
	}
}

func main() {
	const WIDTH = 800
	const HEIGTH = 800
	var MAX_ITERATIONS = 200
	
	var min float64 = -2.84
	var max float64 = 1.0

	var factor float64 = 1
	
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	// Cria window and renderer
	window, renderer, err := sdl.CreateWindowAndRenderer(1400, 900, 0);
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	renderer.SetLogicalSize(WIDTH, HEIGTH)

	count := 0
	for {
		// Esse trecho do algoritmo é responsável por realizar as mudanças nas constantes
		// min e max que são responsáveis por criar o efeito de zoom.
		// As constantes são aleatórias e pode-se mudar o efeito alterando o valor delas
		max -= 0.1*factor
		min += 0.15*factor
		factor *= 0.9349
		MAX_ITERATIONS += 5
		// var file_name string

		if (count > 30) {
			MAX_ITERATIONS = int(float64(MAX_ITERATIONS)*1.02)
		}

		// O PollEvent é necessário para avisar o OS e não deixar a janela irresponsiva
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				break
			}
		}
		renderer.Clear()

		printFractal(WIDTH, HEIGTH, MAX_ITERATIONS, count, min, max, renderer)

		renderer.Present()
		count++
	}
}