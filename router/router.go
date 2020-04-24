package router

import (
	"github.com/FullOfOrange/Devlog-Image/pkg/uploader"
	"net/http"
	"strings"
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

	//query := url.Query()
	//width := query.Get("width")

	imagebyte, err := uploader.GetImage(filename)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	if n, err := w.Write(imagebyte); err != nil || n != len(imagebyte){
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200);
}
