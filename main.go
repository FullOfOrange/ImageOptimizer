package main

import (
	"log"

	"github.com/FullOfOrange/Devlog-Image/pkg/optimizer"
	"github.com/FullOfOrange/Devlog-Image/pkg/uploader"
)

func main() {
	image_file, err := uploader.GetImage("image")
	if err != nil {
		log.Fatal("1we2", err)
	}

	img, err := optimizer.ByteToImage(image_file);
	if err != nil {
		log.Fatal("asdf", err)
	}

	byte, err := img.ResizeWithHeight(300).ImageToPng();
	if err != nil {

		log.Fatal("ㅁㄴㅇㄹ", err)
	}

	uploader.SaveImage(byte)
}
