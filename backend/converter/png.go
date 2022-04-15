package converter

import (
	"errors"
	"github.com/nfnt/resize"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

func ConvertPngToJpeg(pngPath string, destPath string) error {
	pngImgFile, err := os.Open(pngPath)
	if err != nil {
		return errors.New(pngPath + " file not found!")
	}
	defer pngImgFile.Close()

	imgSrc, err := png.Decode(pngImgFile)

	if err != nil {
		return err
	}

	m := resize.Resize(400, 0, imgSrc, resize.Lanczos3)
	newImg := image.NewRGBA(m.Bounds())

	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	draw.Draw(newImg, newImg.Bounds(), m, m.Bounds().Min, draw.Src)

	jpgImgFile, err := os.Create(destPath)

	if err != nil {
		return err
	}

	defer jpgImgFile.Close()

	var opt jpeg.Options
	opt.Quality = 20

	err = jpeg.Encode(jpgImgFile, newImg, &opt)

	if err != nil {
		return err
	}
	err = os.Remove(pngPath)
	if err != nil {
		return err
	}

	return nil
}
