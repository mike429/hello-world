/*
*/

package main

import (
//	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
)

func r_error(w http.ResponseWriter, status int, msg string) {
}

func h_root_get(w http.ResponseWriter, r *http.Request) {
	r_error(w, http.StatusNotImplemented, "GET")
}

func h_root_post(w http.ResponseWriter, r *http.Request) {
/*
	if err := r.ParseForm(); err != nil {
		r_error(w, http.StatusInternalServerError, err.Error())
		goto out
	}
*/

	var (
		reader *multipart.Reader
		part *multipart.Part
		body []byte
		err error
	)
	if reader, err = r.MultipartReader(); err != nil {
		r_error(w, http.StatusInternalServerError, err.Error())
		goto out
	}

	for {
		if part, err = reader.NextPart(); err != nil {
			r_error(w, http.StatusInternalServerError, err.Error())
			goto out
		}

		fmt.Printf("%s\n", part.FormName())
		fmt.Printf("%s\n", part.FileName())

		if body, err = ioutil.ReadAll(part); err != nil {
			r_error(w, http.StatusInternalServerError, err.Error())
			goto out
		}
		fmt.Printf("%s\n", body)
	}

out:
}

func h_root(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(">>>h_root\n")

	switch r.Method {
	case "GET":
		h_root_get(w, r)
	case "POST":
		h_root_post(w, r)
	default:
		r_error(w, http.StatusMethodNotAllowed, r.Method)
	}

	fmt.Printf("<<<h_root\n")
}

func main() {
	log.Printf("...registering handlers")
	http.HandleFunc("/", h_root)

	log.Printf("...serving")
	http.ListenAndServe(":8080", nil)
}


