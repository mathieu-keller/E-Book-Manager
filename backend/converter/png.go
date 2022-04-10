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

	// create image from PNG file
	imgSrc, err := png.Decode(pngImgFile)

	if err != nil {
		return err
	}

	// create a new Image with the same dimension of PNG image
	m := resize.Resize(400, 0, imgSrc, resize.Lanczos3)
	newImg := image.NewRGBA(m.Bounds())

	// we will use white background to replace PNG's transparent background
	// you can change it to whichever color you want with
	// a new color.RGBA{} and use image.NewUniform(color.RGBA{<fill in color>}) function

	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	// paste PNG image OVER to newImage
	draw.Draw(newImg, newImg.Bounds(), m, m.Bounds().Min, draw.Src)

	// create a new Image with the same dimension of PNG image

	// create new out JPEG file
	jpgImgFile, err := os.Create(destPath)

	if err != nil {
		return err
	}

	defer jpgImgFile.Close()

	var opt jpeg.Options
	opt.Quality = 20

	// convert newImage to JPEG encoded byte and save to jpgImgFile
	// with quality = 80
	err = jpeg.Encode(jpgImgFile, newImg, &opt)
	//err = jpeg.Encode(jpgImgFile, newImg, nil) -- use nil if ignore quality options

	if err != nil {
		return err
	}
	err = os.Remove(pngPath)
	if err != nil {
		return err
	}

	return nil
}
