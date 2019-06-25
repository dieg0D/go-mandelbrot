package main

import (
	"math"
)
import "github.com/veandco/go-sdl2/sdl"

func mapValues(value int, inMin, inMax, outMin, outMax float64) float64 {
	return (float64(value)-inMin)*(outMax-outMin)/(inMax-inMin) + outMin
}

func printFractal(WIDTH, HEIGTH, maxIterations int, min, max float64, renderer *sdl.Renderer) {
	// Algoritmo de criação do Mandelbrot
	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGTH; y++ {
			a := mapValues(x, 0, float64(WIDTH), min, max)
			b := mapValues(y, 0, float64(HEIGTH), min, max)

			var ai = a
			var bi = b

			n := 0

			for i := 0; i < maxIterations; i++ {
				var a1 float64
				var b1 float64

				a1 = a*a - b*b
				b1 = 2 * a * b

				a = a1 + ai
				b = b1 + bi

				if (a + b) > 2 {
					break
				}

				n = n + 1
			}

			brigth := mapValues(n, 0, float64(maxIterations), 0, 255)

			// Se a função não divergiu, então o ponto pertence ao conjunto de Mandelbrot
			if n == maxIterations || brigth <= 20 {
				brigth = 0
			}

			// Cores
			red := mapValues(int(brigth*brigth), 0, 255*255, 0, 255)
			green := brigth
			blue := mapValues(int(math.Sqrt(brigth)), 0, math.Sqrt(255), 0, 255)

			renderer.SetDrawColor(uint8(red), uint8(green), uint8(blue), 255)
			renderer.DrawPoint(int32(x), int32(y))
		}
	}
}

func main() {
	const WIDTH = 800
	const HEIGTH = 800
	var maxIterations = 200

	var min = -2.84
	var max = 1.0
	var quit bool

	var factor float64 = 1

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	// Cria window and renderer
	window, renderer, err := sdl.CreateWindowAndRenderer(1200, 700, 0)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	renderer.SetLogicalSize(WIDTH, HEIGTH)

	count := 0
	for !quit {

		// O PollEvent é necessário para avisar o OS e não deixar a janela irresponsiva
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				quit = true
				break
			}
		}

		println("DE FORA", renderer)

		// Esse trecho do algoritmo é responsável por realizar as mudanças nas constantes
		// min e max que são responsáveis por criar o efeito de zoom.
		// As constantes são aleatórias e pode-se mudar o efeito alterando o valor delas
		max -= 0.1 * factor
		min += 0.15 * factor
		factor *= 0.9349
		maxIterations += 5

		if count > 30 {
			maxIterations = int(float64(maxIterations) * 1.02)
		}

		// limpa a teal pra não ficar bugado
		renderer.Clear()

		//Função que desenha o fractal na tela
		printFractal(WIDTH, HEIGTH, maxIterations, min, max, renderer)

		// atualiza a tela
		renderer.Present()
		count++
	}
}
