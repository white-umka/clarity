package ascii

import (
	"net/http"
	"log"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func Converter(img image.Image) [][]string {
	height := img.Bounds().Dy()
	width := img.Bounds().Dx()

	brightness := make([][]float64, height)
	for i := range brightness {
		brightness[i] = make([]float64, width)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixelColor := img.At(x, y) 

			r, g, b, _ := pixelColor.RGBA()

			brightness[y][x] = (0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)) / 655.35
		}
	}

	result := make([][]string, height)
	for i := range result {
		result[i] = make([]string, width)
	}

	for y := range brightness {
		for x := range brightness[y] {
			result[y][x] = GetCharForBrightness(brightness[y][x])
		}
	}
	return result
}

func GetImage(FileURL string) (image.Image, error) {
	resp, err := http.Get(FileURL)
	if err != nil {
		log.Println("ошибка скачивания изображения:", err)
		return nil, err
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body) //
	/*
	 img - image.Image
	 format - png, jpeg и т.д.
	 err - ошибка
	 */

	return img, err
}

func GetCharForBrightness(brightness float64) string { // 1 - черный, 100 - белый
    switch {
    case brightness > 90:
        return "@"  
    case brightness > 80:
        return "0"  
    case brightness > 65:
        return "O"
    case brightness > 45:
        return "o"
    case brightness > 30:
        return "*"
    case brightness > 15:
        return "."
    default:
        return " "  
    }
}