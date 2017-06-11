/*
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var idmap map[string] []string

func r_bad_request(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "%s", msg)
}

func r_method_not_allowed(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(w, "%s", msg)
}

func r_internal_server_error(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "%s", msg)
}

func filter_images(idmap map[string] []string, keys []string) []string {
	fmt.Printf(">>>filter_images\n")

	var ids = []string{}

	if len(keys) == 0 {
		goto out
	}

out:
	fmt.Printf("<<<filter_images\n")
	return ids
}

func parse_ids(values map[string][]string) []string {
	var ids, arr []string
	var ok bool

	ids = []string{}

	if arr, ok = values["key"]; !ok {
		goto out
	}

	for _, list := range arr {
		tags := strings.Split(list, ",")
		ids = append(ids, tags...)
	}

out:
	return ids
}

/*
GET /?key=tag,...&key=tag,...
*/
func h_root_get(w http.ResponseWriter, r *http.Request, ids []string) {
	fmt.Printf(">>>h_root_get\n")

	imgs := filter_images(idmap, ids)
	for _, img := range imgs {
		fmt.Printf("%2d: %s\n", img)
	}

	fmt.Printf("<<<h_root_get\n")
}

/*
POST /?id=id&key=tag&key=tag...
TODO:
1.  Handle initial upload (when there is no image id)
*/
func h_root_post(w http.ResponseWriter, r *http.Request, ids []string) {
	fmt.Printf(">>>h_root_post\n")
	fmt.Printf("<<<h_root_post\n")
}

func h_root(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(">>>h_root\n")

	var ids []string

	if err := r.ParseForm(); err != nil {
		r_internal_server_error(w, err.Error())
		goto out
	}

	ids = parse_ids(r.Form)
/*
TODO:
What does empty ids mean?  Everything or nothing?
*/

	switch r.Method {
	case "GET":
		h_root_get(w, r, ids)
	case "POST":
		h_root_post(w, r, ids)
	default:
		r_method_not_allowed(w, r.Method)
	}
out:
	fmt.Printf("<<<h_root\n")
}

func bogus_data() {
	log.Printf(">>>bogus data\n")

	idmap["u1"] = []string{"i1"}
	idmap["u2"] = []string{"i2"}
	idmap["u3"] = []string{"i3"}

	idmap["s12"] = []string{"i1", "i2"}
	idmap["s23"] = []string{"i2", "i3"}
	idmap["s13"] = []string{"i1", "i3"}

	idmap["s123"] = []string{"i1", "i2", "i3"}

	log.Printf("<<<bogus data\n")
}

func main() {

	idmap = make(map[string] []string)

	bogus_data()

	log.Printf("...registering handlers")
	http.HandleFunc("/", h_root)

	log.Printf("...serving")
	http.ListenAndServe(":8080", nil)
}


