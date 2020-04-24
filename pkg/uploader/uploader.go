package uploader

import (
	"fmt"
	"os"
	"io/ioutil"

	"github.com/google/uuid"
)

func GetImage(uuid string) ([]byte, error){
	f,err := os.Open(fmt.Sprintf("%s.png", uuid))
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
	f, err := os.Create(fmt.Sprintf("%s.png", name))
	if err != nil {
		return name, err
	}
	// 만약 쓰기 실패가 발생하더라도 추후 이미지 사용률에 따라 제거될 것이니
	// 신경쓰지 않고 이미지를 지우는 로직은 작성은 일단 안해놓
	n, err := f.Write(image)
	if err != nil {
		return name, err
	}
	if  n != len(image){
		return name, fmt.Errorf("writing error")
	}

	defer f.Close();
	return name, nil;
}
