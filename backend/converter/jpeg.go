package converter

import (
	"errors"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"os"
)

func CompressImageResource(jpegPath string) error {
	file, err := os.Open(jpegPath)
	if err != nil {
		return errors.New(jpegPath + " file not found!")
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	// Set the expected size that you want:
	newImg := resize.Resize(400, 0, img, resize.Lanczos3)

	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}
	err = os.Remove(jpegPath)
	if err != nil {
		return err
	}
	jpgImgFile, err := os.Create(jpegPath)
	if err != nil {
		return err
	}
	defer jpgImgFile.Close()
	err = jpeg.Encode(jpgImgFile, newImg, &jpeg.Options{Quality: 20})
	if err != nil {
		return err
	}
	return nil
}
