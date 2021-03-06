package optimizer

import (
	"bytes"
	"image"
	"image/png"
)

func encode(src *image.Image, buf *bytes.Buffer, done chan error) {
	if err := png.Encode(buf, *src); err != nil {
		done <- err
		close(done)
		return
	}
	done <- nil
	close(done)
}

func (m *Image) ImageToPNGByte() ([]byte, error) {
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