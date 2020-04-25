package router

import (
	"fmt"
	"github.com/FullOfOrange/ImageOptimizer/pkg/optimizer"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/FullOfOrange/ImageOptimizer/pkg/uploader"
)

func InitRouter() *http.ServeMux{
	mux := http.NewServeMux()
	mux.HandleFunc("/", imageHandler)
	return mux
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
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
					imagebyte, err = image.Resize(width, height).ImageToJPEGByte()
				} else if width != 0 {
					imagebyte, err = image.ResizeWithWidth(width).ImageToJPEGByte()
				} else {
					imagebyte, err = image.ResizeWithHeight(height).ImageToJPEGByte()
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
	} else if r.Method == http.MethodPost {

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