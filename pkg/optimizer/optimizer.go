package optimizer

import (
	"image"
	_ "image/gif"
	_ "image/png"

	"github.com/disintegration/imaging"
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