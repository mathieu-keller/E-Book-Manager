package converter

import (
	"errors"
	"github.com/nfnt/resize"
	"image/gif"
	"image/jpeg"
	"os"
)

func ConvertGifToJpeg(pngPath string, destPath string) error {
	pngImgFile, err := os.Open(pngPath)
	if err != nil {
		return errors.New(pngPath + " file not found!")
	}
	defer pngImgFile.Close()

	// create image from PNG file
	imgSrc, err := gif.Decode(pngImgFile)

	if err != nil {
		return err
	}

	// create a new Image with the same dimension of PNG image
	newImg := resize.Resize(400, 0, imgSrc, resize.Lanczos3)

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
