package converter

import (
	"bytes"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
)

const Quality = 50

func CompressImageResource(byteFile []byte) ([]byte, error) {
	file := bytes.NewReader(byteFile)
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	newImg := resize.Resize(270, 0, img, resize.Lanczos3)
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, newImg, &jpeg.Options{Quality: Quality})
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
