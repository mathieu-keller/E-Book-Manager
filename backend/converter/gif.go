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

	imgSrc, err := gif.Decode(pngImgFile)

	if err != nil {
		return err
	}

	newImg := resize.Resize(270, 0, imgSrc, resize.Lanczos3)
	jpgImgFile, err := os.Create(destPath)

	if err != nil {
		return err
	}

	defer jpgImgFile.Close()

	var opt jpeg.Options
	opt.Quality = Quality

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
