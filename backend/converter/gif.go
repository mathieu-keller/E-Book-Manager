package converter

import (
	"bytes"
	"github.com/nfnt/resize"
	"image/gif"
	"image/jpeg"
)

func ConvertGifToJpeg(byteFile []byte) ([]byte, error) {
	file := bytes.NewReader(byteFile)

	imgSrc, err := gif.Decode(file)

	if err != nil {
		return nil, err
	}

	newImg := resize.Resize(270, 0, imgSrc, resize.Lanczos3)

	var opt jpeg.Options
	opt.Quality = Quality
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, newImg, &opt)

	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
