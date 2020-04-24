package optimizer

import (
	"bytes"
	"image"
	"image/png"
	_ "image/jpeg"
	_ "image/gif"

	"github.com/disintegration/imaging"
)

type Image struct {
	src *image.Image
	conf *image.Config
}

var pngEncoder = png.Encoder{CompressionLevel: -3}



func (m *Image) Resize(width int, height int) *Image {
	resize_image := imaging.Resize(*m.src, width, height, imaging.Lanczos);

	*m.src = resize_image
	return m
}

func (m *Image) ResizeWithHeight (height int) *Image {
	p := height / m.conf.Height
	m.Resize(m.conf.Width * p, height)
	return m
}

func (m *Image) ResizeWithWidth (width int) *Image {
	p := width / m.conf.Width
	m.Resize(width, m.conf.Height * p)
	return m
}

func (m *Image) ImageToPngByte() ([]byte, error) {
	buf := new(bytes.Buffer);

	if err := pngEncoder.Encode(buf, *m.src); err != nil {
		return nil, err;
	}
	return buf.Bytes(), nil
}

func ByteToImage(byte []byte) (*Image, error) {
	//contentType := http.DetectContentType(byte)

	reader := bytes.NewReader(byte)
	img, _, err := image.Decode(reader);
	if err != nil {
		return nil, err
	}
	reader.Reset(byte);
	config, _, err := image.DecodeConfig(reader)
	if err != nil {
		return nil, err
	}

	return &Image{&img, &config}, nil;
}