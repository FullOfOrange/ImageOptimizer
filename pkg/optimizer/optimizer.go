package optimizer

import (
	"bytes"
	"github.com/disintegration/imaging"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
)

type Image struct {
	src *image.Image
	conf *image.Config
}

func resize(m *Image, width int, height int, done chan bool) {
	*m.src = imaging.Resize(*m.src, width, height, imaging.Lanczos);
	done <- true
	close(done)
}

func (m *Image) Resize(width int, height int) *Image {
	done := make(chan bool)
	go resize(m, width, height, done);

	<- done
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

func encode(src *image.Image, buf *bytes.Buffer, done chan error) {
	if err := jpeg.Encode(buf, *src, nil); err != nil {
		done <- err
		close(done)
		return
	}
	done <- nil
	close(done)
}

func (m *Image) ImageToJPEGByte() ([]byte, error) {
	buf := new(bytes.Buffer);
	done := make(chan error)

	go encode(m.src, buf, done)

	if err :=<- done; err != nil {
		return nil, err
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