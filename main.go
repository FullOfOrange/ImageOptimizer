package main

import (
	"github.com/FullOfOrange/Devlog-Image/pkg/optimizer"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	image_file, err := ioutil.ReadFile("./image.png")
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

	f, err := os.Create("image2.png");
	f.Write(byte)
}
