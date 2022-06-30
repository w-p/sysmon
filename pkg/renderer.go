package pkg

import (
	"bufio"
	"bytes"
	"math"

	"github.com/fogleman/gg"
)

const (
	width   = 8.0
	padding = 6.0
)

var scaler ScaleFn
var keys = []string{"cpu", "ram"}
var palette = map[string]string{
	"cpu": "#72f542",
	"ram": "#b833ff",
}

// Render produces the icon based on stats values
func Render(stats Stats) []byte {
	if scaler == nil {
		scaler = NewScale(
			Bounds{min: 0.0, max: 100.0},
			Bounds{min: 0.0, max: 24.0},
		)
	}
	x := 0.0
	ctx := gg.NewContext(24, 24)

	for _, key := range keys {
		color := palette[key]
		value := scaler(stats[key])
		pretty := math.Round(value*100) / 100
		ctx.DrawRectangle(x, 24, width, pretty*-1)
		ctx.SetHexColor(color)
		ctx.Fill()
		x += width + padding
	}

	var bytes bytes.Buffer
	w := bufio.NewWriter(&bytes)
	ctx.EncodePNG(w)
	err := w.Flush()
	if err != nil {
		panic("failed to flush writer")
	}

	return bytes.Bytes()
}
