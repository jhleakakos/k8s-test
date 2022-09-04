package main

import (
	"io"
	"net/http"
)

func home(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, `<p style="text-align: center; padding-top: 10%;">Aren\'t you a good k8s: <b style="font-size: 2rem;">v1 from golang</b></p>`)
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8000", nil)

}