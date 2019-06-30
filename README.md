
Para instalar o sdl : go get -v github.com/veandco/go-sdl2/{sdl,img,mix,ttf}

## Iniciando o Projeto
O arquivo principal do projeto é o `mandelbrot.go`. Compile o programa usando:

```
    go build mandelbrot.go
```
E depois execute com:

```
    ./mandelbrot
```

## Processamento Paralelo
Para compilar o arquivo que explora o paralelismo em Go rode o seguinte comando:

```
    go build paralell_mandelbrot.go
```

Esse programa aceita argumentos de linha de comando no seguinte formato:

```
    ./paralell_mandelbrot TIPO_DE_EXECUÇÃO NUMERO_DE_IMAGENS
```
Onde `TIPO_DE_EXECUÇÃO` deve ser a string `"sequencial"`, para processamento sequencial, ou
`"paralelo"`, para processamento paralelo, e o número de imagens corresponde a quantidade de imagens que se quer
gerar na pasta `/png`.

Ao final da execução, o programa irá mostrar o tempo de execução do programa para que se possa comparar.