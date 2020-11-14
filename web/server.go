package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"../steg"
)

func hideRoute(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		var buf = new(bytes.Buffer)
		err := req.ParseMultipartForm(1 << 25)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		fmt.Println(req.Header.Get("Content-Type"))
		file, header, err := req.FormFile("img")
		if err != nil {
			w.WriteHeader(400)
			return
		}
		defer file.Close()
		name := header.Filename
		fmt.Println(name)
		io.Copy(buf, file)
		// fmt.Println("buf: ", buf.String())
		m := req.FormValue("m")
		fmt.Println(m)
		newBuf, err := steg.EncodeFromFile(buf, name, m)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		w.Header().Set("Content-Type", "image/jpeg")
		w.WriteHeader(200)
		w.Write(newBuf.Bytes())
		// body, err := ioutil.ReadAll(req.Body)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(string(body))
	} else {
		fmt.Fprintf(w, "test route\n")
	}
}

func revealRoute(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {

	} else {
		fmt.Fprintf(w, "decode route\n")
	}
}

func main() {
	fs := http.FileServer(http.Dir("./public"))

	http.Handle("/", fs)
	http.HandleFunc("/hide", hideRoute)
	http.HandleFunc("/reveal", revealRoute)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
