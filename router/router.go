package router

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/FullOfOrange/Devlog-Image/pkg/optimizer"
	"github.com/FullOfOrange/Devlog-Image/pkg/uploader"
)

func InitRouter() *http.ServeMux{
	mux := http.NewServeMux()
	mux.HandleFunc("/", imageDownloadHandler)
	return mux
}

func imageDownloadHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		w.WriteHeader(404)
		return
	}

	url := r.URL
	filename := strings.TrimLeft(url.Path,"/");

	query := url.Query()

	imagebyte, err := uploader.GetImage(filename)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	if width, err := strconv.Atoi(query.Get("width")); err == nil {
		image, err := optimizer.ByteToImage(imagebyte);
		if err != nil {
			w.WriteHeader(500)
			return
		}
		imagebyte, err = image.ResizeWithWidth(width).ImageToJPEGByte()
		if err != nil {
			w.WriteHeader(500)
			return
		}
	}

	w.Header().Add("Content-Type","image/jpeg")
	w.WriteHeader(200);
	if n, err := w.Write(imagebyte); err != nil || n != len(imagebyte){
		w.WriteHeader(500)
		return
	}
}
