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
		var imgBuf = new(bytes.Buffer)
		var fileBuf = new(bytes.Buffer)
		err := req.ParseMultipartForm(1 << 25)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		// fmt.Println(req.Header.Get("Content-Type"))
		imgFile, _, err := req.FormFile("img")
		if err != nil {
			w.WriteHeader(400)
			return
		}
		defer imgFile.Close()
		fileFile, fileHeader, _ := req.FormFile("file")
		var fileName string
		if fileFile != nil {
			fileName = fileHeader.Filename
			io.Copy(fileBuf, fileFile)
			defer fileFile.Close()
		}
		io.Copy(imgBuf, imgFile)
		// fmt.Println("buf: ", buf.String())
		m := req.FormValue("m")
		fmt.Println(m)
		newImgBuf, err := steg.EncodeFromFile(imgBuf, m, fileBuf.Bytes(), fileName)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		w.Header().Set("Content-Type", "image/jpeg")
		w.WriteHeader(200)
		w.Write(newImgBuf.Bytes())
		// body, err := ioutil.ReadAll(req.Body)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(string(body))
	} else {
		fmt.Fprintf(w, "hide route\n")
	}
}

func revealRoute(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		err := req.ParseMultipartForm(1 << 25)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		imgFile, _, err := req.FormFile("img")
		if err != nil {
			w.WriteHeader(400)
			return
		}
		defer imgFile.Close()
		var imgBuf = new(bytes.Buffer)
		io.Copy(imgBuf, imgFile)
		message, filename, err := steg.DecodeFromFile(imgBuf, "public/uploads")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("message: " + message)
		fmt.Println("filename: " + filename)
		response := "<html><body><h1>Message</h1><div>"
		response += message
		response += "</div>"
		if filename != "" {
			response += "<div><a href=\""
			response += "/uploads/"
			response += filename + "\">"
			response += "File</a></div>"
		}
		response += "</body></html>"
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(200)
		w.Write([]byte(response))
		fmt.Println(response)
	} else {
		fmt.Fprintf(w, "reveal route\n")
	}
}

func main() {
	fs := http.FileServer(http.Dir("./public"))

	http.Handle("/", fs)
	http.HandleFunc("/hide", hideRoute)
	http.HandleFunc("/reveal", revealRoute)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
