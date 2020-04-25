package main

import (
	"net/http"

	"github.com/FullOfOrange/ImageOptimizer/router"
)

func main() {
	mux := router.InitRouter()

	http.ListenAndServe(":8080", mux)

}
