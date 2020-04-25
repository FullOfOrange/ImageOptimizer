package router

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/FullOfOrange/ImageOptimizer/pkg/optimizer"
	"github.com/FullOfOrange/ImageOptimizer/pkg/uploader"
)

func InitRouter() *http.ServeMux{
	mux := http.NewServeMux()
	mux.HandleFunc("/", imageHandler)
	return mux
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	// GET 으로 들어온 파일 다운로드 요청
	if r.Method == http.MethodGet {
		url := r.URL
		filename := strings.TrimLeft(url.Path, "/");
		width, height := getImageOptSize(url.Query())
		str := filename +"?"+getImageOptQueryString(width, height)

		imagebyte, err := uploader.GetImage(str)
		if err != nil {
			imagebyte, err = uploader.GetImage(filename)
			if err != nil {
				w.WriteHeader(404)
				return
			}
			if width != 0 || height != 0 {
				image, err := optimizer.ByteToImage(imagebyte);
				if err != nil {
					w.WriteHeader(500)
					return
				}
				if width != 0 && height != 0 {
					imagebyte, err = image.Resize(width, height).ImageToPNGByte()
				} else if width != 0 {
					imagebyte, err = image.ResizeWithWidth(width).ImageToPNGByte()
				} else {
					imagebyte, err = image.ResizeWithHeight(height).ImageToPNGByte()
				}
				if err != nil {
					w.WriteHeader(500)
					return
				}
			}
			uploader.SaveImage(imagebyte, str)
		}
		w.Header().Add("Content-Type", "image/jpeg")
		w.WriteHeader(200);
		if n, err := w.Write(imagebyte); err != nil || n != len(imagebyte) {
			w.WriteHeader(500)
			return
		}

	// POST 로 들어온 파일 업로드 요청임
	} else if r.Method == http.MethodPost {

	// Delete 로 들어온 파일 삭제 요청임
	// 이곳에서는 캐시와 optimizing 된 모든 파일 또한 삭제되어야함.
	} else if r.Method == http.MethodDelete{
		
	} else {
		w.WriteHeader(404)
		return
	}
}

func getImageOptSize(query url.Values) (width int, height int) {
	var err error
	width, err = strconv.Atoi(query.Get("width"));
	if err != nil {
		width = 0;
	}
	height, err = strconv.Atoi(query.Get("height"));
	if err != nil {
		height = 0;
	}
	return
}

func getImageOptQueryString(width int, height int) (result string){
	if width != 0 {
		result += fmt.Sprintf("width=%d", width)
	}
	if height != 0 {
		result += fmt.Sprintf("height=%d", height)
	}
	return
}