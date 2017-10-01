// Lissajous generstes GIF animation of random Lissajous figures.
package main

import (
	"image/color"
	"io"
	"math/rand"
	"image/gif"
	"image"
	"math"
	"os"
)

var paletee = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in paletee
	blackIndex = 1 // second color in paletee
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycle = 5 // number of complete x oscillator revolutions
		res = 0.001 //angular resolution
		size = 100 // image canavas covers [-size...+size]
		nframes = 64 // number of animation frames
		delay = 8 // delay between frames in 10ms unit
	)
	freq := rand.Float64() * 0.3 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, paletee)
		for t := 0.0; i < cycle * 2 * math.Pi; i += res {
			x := math.Sin(t)
			y := math.Sin(t * freq + phase)
			img.SetColorIndex(size+int(x * size + 0.5), size+int(y * size + 0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
