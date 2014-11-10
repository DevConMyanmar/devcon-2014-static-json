package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type data []map[string]string

func returnJSON(w http.ResponseWriter, r *http.Request) {
	url := strings.Replace(r.URL.String(), "/", "", 1)

	jsonFile, err := ioutil.ReadFile("./" + url + ".json")

	var p data

	_ = json.Unmarshal(jsonFile, &p)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(404) // HTTP 500
		errorMsg := map[string]error{"msg": err}
		mapB, _ := json.Marshal(errorMsg)
		w.Write(mapB)
	} else {
		w.Write(jsonFile)
	}
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	http.HandleFunc("/", returnJSON)
}
