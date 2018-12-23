package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime"
	"os"
	"strings"

	"mime/multipart"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Listening on port 4141")

	router := mux.NewRouter()
	router.HandleFunc("/v1/upload", MimeUpload).Methods("POST")
	log.Fatal(http.ListenAndServe(":4141", router))
}

func MimeUpload(w http.ResponseWriter, r *http.Request) {

	mediaType, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		log.Fatal(err)
	}
	if strings.HasPrefix(mediaType, "multipart/") {
		mr := multipart.NewReader(r.Body, params["boundary"])
		for {
			p, err := mr.NextPart()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatal(err)
			}

			//In Bytes Type
			// slurp, err := ioutil.ReadAll(p)
			// if err != nil {
			// 	log.Fatal(err)
			// }

			fmt.Printf("%s:\n", p.Header.Get("Content-Disposition"))

			s := strings.Replace(p.Header.Get("Content-Disposition"), " ", "", -1)
			s = strings.Replace(s, "\"", "", -1)
			ss := strings.Split(s, ";")
			var m map[string]string
			m = make(map[string]string)
			for _, pair := range ss[1:] {
				z := strings.Split(pair, "=")
				m[z[0]] = z[1]
			}

			if len(m["filename"]) > 0 {

				//Save to File
				file, err := os.Create(m["filename"])
				if err != nil {
					fmt.Println(err)
				}
				defer file.Close()

				_, err = io.Copy(file, p)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("Saved file: %+v\n", m["filename"])
				}

			}
		}
	}

	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	fmt.Println(strings.Join(request, "\n"))
	json.NewEncoder(w).Encode(request)
}
