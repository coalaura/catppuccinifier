package main

import (
	"image"
	"image/png"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	_ "image/gif"
	_ "image/jpeg"

	_ "golang.org/x/image/webp"

	"github.com/coalaura/plain"
)

var log = plain.New()

func main() {
	if len(os.Args) < 2 {
		log.Errorln("Usage: catppuccin <input> [output]")

		os.Exit(1)
	}

	input := os.Args[1]

	var output string

	if len(os.Args) > 2 {
		output = os.Args[2]
	} else {
		ext := filepath.Ext(input)
		base := strings.TrimSuffix(input, ext)

		output = base + "_catppuccin.png"
	}

	log.Println("Reading input...")

	file, err := os.Open(input)
	log.MustFail(err)

	defer file.Close()

	img, _, err := image.Decode(file)
	log.MustFail(err)

	bounds := img.Bounds()
	out := image.NewRGBA(bounds)

	log.Println("Processing image...")

	var wg sync.WaitGroup

	rowChan := make(chan int, bounds.Dy())

	for range runtime.NumCPU() {
		wg.Go(func() {
			for y := range rowChan {
				for x := bounds.Min.X; x < bounds.Max.X; x++ {
					out.Set(x, y, mapToPalette(img.At(x, y)))
				}
			}
		})
	}

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		rowChan <- y
	}

	close(rowChan)
	wg.Wait()

	log.Printf("Writing output to %s...\n", output)

	outFile, err := os.Create(output)
	log.MustFail(err)

	defer outFile.Close()

	err = png.Encode(outFile, out)
	log.MustFail(err)

	log.Println("Done!")
}
