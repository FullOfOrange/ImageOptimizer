package uploader

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

func GetImage(name string) ([]byte, error){
	var byte []byte
	f, err := os.Open(fmt.Sprintf("./images/%s", name))
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

func SaveImage(image []byte, name string) (string, error) {
	if name == "" {
		name = uuid.New().String();
	}
	filename := fmt.Sprintf("./images/%s", name)

	f, err := os.Create(filename)
	if err != nil {
		return name, err
	}

	if n, err := f.Write(image); err != nil || n != len(image) {
		if err = os.Remove(filename); err != nil {
			return name, fmt.Errorf("writing error but can't remove")
		}
		return name, fmt.Errorf("writing error")
	}

	defer f.Close();
	return name, nil;
}
