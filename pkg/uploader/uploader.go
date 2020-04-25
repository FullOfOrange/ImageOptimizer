package uploader

import (
	"fmt"
	"github.com/FullOfOrange/ImageOptimizer/pkg/cache"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

func GetImage(uuid string) ([]byte, error){
	var byte []byte
	if !cache.CheckCachedImage(uuid) {
		f, err := os.Open(fmt.Sprintf("./images/%s", uuid))
		if err != nil {
			return nil, err
		}

		byte, err = ioutil.ReadAll(f);
		if err != nil {
			return nil, err
		}

		cache.CachingImage(uuid, byte)

		defer f.Close();
	} else {
		byte = cache.GetCachedImage(uuid)
	}

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
