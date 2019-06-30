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
	"time"
	"log"
	"sync"
)

func map_values(value int, in_min, in_max, out_min, out_max float64) float64 {
	return (float64(value) - in_min) * (out_max - out_min) / (in_max - in_min) + out_min
}

func parallelFractal(WIDTH, HEIGTH, maxIterations, count int, min, max float64, wg *sync.WaitGroup) {
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
	defer wg.Done()
}

func sequentialFractal(WIDTH, HEIGTH, maxIterations, count int, min, max float64) {
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
	// Primeiro argumento: usar processamento paralelo ou não
	// Segundo argumento: número de imanges a serem geradas
	// args := os.Args[1:]
	execution_type := os.Args[1]
	// n_images := os.Args[2]

	n_images, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}

	const WIDTH = 800
	const HEIGTH = 800
	var MAX_ITERATIONS = 200
	
	var min float64 = -2.84
	var max float64 = 1.0

	var factor float64 = 1

	// Criando WaitGroup para sincronizar as tarefas paralelas
	var wg sync.WaitGroup
	wg.Add(n_images)

	count := 0
	start := time.Now()
	switch execution_type {
	case "sequencial":
		for {
			// Esse trecho do algoritmo é responsável por realizar as mudanças nas constantes
			// min e max que são responsáveis por criar o efeito de zoom.
			// As constantes são aleatórias e pode-se mudar o efeito alterando o valor delas
			max -= 0.1*factor
			min += 0.15*factor
			factor *= 0.9349
			MAX_ITERATIONS += 5
	
			if (count > n_images) {
				elapsed := time.Since(start)
				log.Printf("Execution took %s", elapsed)
				os.Exit(0)
			}
	
			sequentialFractal(WIDTH, HEIGTH, MAX_ITERATIONS, count, min, max)
	
			count++
		}
	case "paralelo":
		for {
			// Esse trecho do algoritmo é responsável por realizar as mudanças nas constantes
			// min e max que são responsáveis por criar o efeito de zoom.
			// As constantes são aleatórias e pode-se mudar o efeito alterando o valor delas
			max -= 0.1*factor
			min += 0.15*factor
			factor *= 0.9349
			MAX_ITERATIONS += 5
	
			if (count > n_images) {
				wg.Wait()
				elapsed := time.Since(start)
				log.Printf("Execution took %s", elapsed)
				os.Exit(0)
			}
	
			go parallelFractal(WIDTH, HEIGTH, MAX_ITERATIONS, count, min, max, &wg)
	
			count++
		}
	}
}