package uploader

import (
	"fmt"
	"os"
	"io/ioutil"

	"github.com/google/uuid"
)

func GetImage(uuid string) ([]byte, error){
	f, err := os.Open(fmt.Sprintf("./images/%s.png", uuid))
	if err != nil {
		return nil, err
	}

	byte, err := ioutil.ReadAll(f);
	if err != nil {
		return nil, err
	}

	defer f.Close();
	return byte, nil;
}

func SaveImage(image []byte) (string, error) {
	name := uuid.New().String();
	filename := fmt.Sprintf("./images/%s.png", name)

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
