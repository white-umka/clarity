package ascii

import (
	"bytes"
	_ "embed"
	"strings"

	"github.com/fogleman/gg"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"golang.org/x/image/font/opentype"
)

//go:embed JetBrainsMono-Regular.ttf
var fontBytes []byte

func Render(asciiArt [][]string) (tgbotapi.FileBytes, error) {
	if len(asciiArt) == 0 || len(asciiArt[0]) == 0 {
		return tgbotapi.FileBytes{}, nil
	}

	parsedFont, err := opentype.Parse(fontBytes)
	if err != nil {
		return tgbotapi.FileBytes{}, err
	}

	const fontSize = 12.0
	face, err := opentype.NewFace(parsedFont, &opentype.FaceOptions{
		Size: fontSize,
		DPI:  72,
	})
	if err != nil {
		return tgbotapi.FileBytes{}, err
	}
	defer face.Close()

	const charWidth = 7.2
	const lineHeight = 12.0

	rows := len(asciiArt)
	cols := len(asciiArt[0])

	imgWidth := int(float64(cols)*charWidth) + 20
	imgHeight := int(float64(rows)*lineHeight) + 20

	dc := gg.NewContext(imgWidth, imgHeight)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	dc.SetFontFace(face)
	dc.SetRGB(0, 0, 0)

	for y, row := range asciiArt {
		line := strings.Join(row, "")
		dc.DrawString(line, 10, 10+float64(y+1)*lineHeight)
	}

	var buf bytes.Buffer
	if err := dc.EncodePNG(&buf); err != nil {
		return tgbotapi.FileBytes{}, err
	}

	return tgbotapi.FileBytes{
		Name:  "ascii_art.png",
		Bytes: buf.Bytes(),
	}, nil
}
