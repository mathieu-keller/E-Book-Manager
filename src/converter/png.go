package converter

import (
	"bytes"
	"github.com/nfnt/resize"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
)

func ConvertPngToJpeg(byteFile []byte) ([]byte, error) {
	file := bytes.NewReader(byteFile)
	imgSrc, err := png.Decode(file)

	if err != nil {
		return nil, err
	}

	m := resize.Resize(270, 0, imgSrc, resize.Lanczos3)
	newImg := image.NewRGBA(m.Bounds())

	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	draw.Draw(newImg, newImg.Bounds(), m, m.Bounds().Min, draw.Src)

	var opt jpeg.Options
	opt.Quality = Quality
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, newImg, &opt)

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
