package uploader

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/uuid"
)

var IMAGE_DIR = "./images"
var ORIGIN = "/ori"
var OPTIMIZED = "/opt"

func getImagePath(name string, isOrigin bool) string {
	var dir = ORIGIN
	if !isOrigin {
		dir = OPTIMIZED
	}
	return IMAGE_DIR + dir + "/" + name
}

func GetImage(name string, isOrigin bool) ([]byte, error){
	var byte []byte
	str := getImagePath(name, isOrigin)
	log.Print(str)
	f, err := os.Open(getImagePath(name, isOrigin))
	if err != nil {
		return nil, err
	}

	byte, err = ioutil.ReadAll(f);
	if err != nil {
		return nil, err
	}

	defer f.Close();
	return byte, nil;
}

func SaveImage(image []byte, name string, isOrigin bool) (string, error) {
	var err error
	var f *os.File

	if name == "" {
		name = uuid.New().String();
	}
	filename := getImagePath(name, isOrigin)

	if f = CheckImageExist(name, isOrigin); f == nil {
		if f, err = os.Create(filename); err != nil {
			return name, err
		}
	}

	if n, err := f.Write(image); err != nil || n != len(image) {
		if err = os.Remove(filename); err != nil {
			return name, fmt.Errorf("writing error but can't remove")
		}
		return name, fmt.Errorf("writing error")
	}

	defer f.Close()
	return name, nil
}

func CheckImageExist(name string, isOrigin bool) *os.File {
	if f, err := os.Open(getImagePath(name, isOrigin)); err != nil {
		return f
	}
	return nil
}
