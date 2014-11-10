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

	jsonFile, _ := ioutil.ReadFile("./" + url + ".json")

	var p data

	_ = json.Unmarshal(jsonFile, &p)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonFile)
}

func init() {
	http.HandleFunc("/", returnJSON)

	// err := http.ListenAndServe(":9090", nil)

	//	if err != nil {
	//		log.Fatal("ListenAndServe: ", err)
	//	}
}
