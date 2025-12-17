package main

import (
	"image/color"
	"math"
	"sync"

	"github.com/lucasb-eyer/go-colorful"
)

var (
	palette    []colorful.Color
	colorCache sync.Map
)

func init() {
	hexes := []string{
		"#dc8a78",
		"#dd7878",
		"#ea76cb",
		"#8839ef",
		"#d20f39",
		"#e64553",
		"#fe640b",
		"#df8e1d",
		"#40a02b",
		"#179299",
		"#04a5e5",
		"#209fb5",
		"#1e66f5",
		"#7287fd",
		"#4c4f69",
		"#5c5f77",
		"#6c6f85",
		"#7c7f93",
		"#8c8fa1",
		"#9ca0b0",
		"#acb0be",
		"#bcc0cc",
		"#ccd0da",
		"#eff1f5",
		"#e6e9ef",
		"#dce0e8",
		"#5c5f77",
		"#6c6f85",
		"#d20f39",
		"#de293e",
		"#40a02b",
		"#49af3d",
		"#df8e1d",
		"#eea02d",
		"#1e66f5",
		"#456eff",
		"#ea76cb",
		"#fe85d8",
		"#179299",
		"#2d9fa8",
		"#acb0be",
		"#bcc0cc",
		"#f2d5cf",
		"#eebebe",
		"#f4b8e4",
		"#ca9ee6",
		"#e78284",
		"#ea999c",
		"#ef9f76",
		"#e5c890",
		"#a6d189",
		"#81c8be",
		"#99d1db",
		"#85c1dc",
		"#8caaee",
		"#babbf1",
		"#c6d0f5",
		"#b5bfe2",
		"#a5adce",
		"#949cbb",
		"#838ba7",
		"#737994",
		"#626880",
		"#51576d",
		"#414559",
		"#303446",
		"#292c3c",
		"#232634",
		"#51576d",
		"#626880",
		"#e78284",
		"#e67172",
		"#a6d189",
		"#8ec772",
		"#e5c890",
		"#d9ba73",
		"#8caaee",
		"#7b9ef0",
		"#f4b8e4",
		"#f2a4db",
		"#81c8be",
		"#5abfb5",
		"#a5adce",
		"#b5bfe2",
		"#f4dbd6",
		"#f0c6c6",
		"#f5bde6",
		"#c6a0f6",
		"#ed8796",
		"#ee99a0",
		"#f5a97f",
		"#eed49f",
		"#a6da95",
		"#8bd5ca",
		"#91d7e3",
		"#7dc4e4",
		"#8aadf4",
		"#b7bdf8",
		"#cad3f5",
		"#b8c0e0",
		"#a5adcb",
		"#939ab7",
		"#8087a2",
		"#6e738d",
		"#5b6078",
		"#494d64",
		"#363a4f",
		"#24273a",
		"#1e2030",
		"#181926",
		"#494d64",
		"#5b6078",
		"#ed8796",
		"#ec7486",
		"#a6da95",
		"#8ccf7f",
		"#eed49f",
		"#e1c682",
		"#8aadf4",
		"#78a1f6",
		"#f5bde6",
		"#f2a9dd",
		"#8bd5ca",
		"#63cbc0",
		"#a5adcb",
		"#b8c0e0",
		"#f5e0dc",
		"#f2cdcd",
		"#f5c2e7",
		"#cba6f7",
		"#f38ba8",
		"#eba0ac",
		"#fab387",
		"#f9e2af",
		"#a6e3a1",
		"#94e2d5",
		"#89dceb",
		"#74c7ec",
		"#89b4fa",
		"#b4befe",
		"#cdd6f4",
		"#bac2de",
		"#a6adc8",
		"#9399b2",
		"#7f849c",
		"#6c7086",
		"#585b70",
		"#45475a",
		"#313244",
		"#1e1e2e",
		"#181825",
		"#11111b",
		"#45475a",
		"#585b70",
		"#f38ba8",
		"#f37799",
		"#a6e3a1",
		"#89d88b",
		"#f9e2af",
		"#ebd391",
		"#89b4fa",
		"#74a8fc",
		"#f5c2e7",
		"#f2aede",
		"#94e2d5",
		"#6bd7ca",
		"#a6adc8",
		"#bac2de",
	}

	for _, h := range hexes {
		c, _ := colorful.Hex(h)

		palette = append(palette, c)
	}
}

func mapToPalette(src color.Color) color.Color {
	r, g, b, a := src.RGBA()

	if a == 0 {
		return color.RGBA{0, 0, 0, 0}
	}

	key := [4]uint32{r, g, b, a}

	if val, ok := colorCache.Load(key); ok {
		return val.(color.Color)
	}

	target, _ := colorful.MakeColor(src)

	var best colorful.Color

	minDist := math.MaxFloat64

	for _, p := range palette {
		d := target.DistanceLab(p)

		if d < minDist {
			minDist = d
			best = p
		}
	}

	br, bg, bb := best.RGB255()

	final := color.RGBA{br, bg, bb, uint8(a >> 8)}

	colorCache.Store(key, final)

	return final
}
