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

	jsonFile, err := ioutil.ReadFile(url + ".json")

	var p data

	_ = json.Unmarshal(jsonFile, &p)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		errorMsg := map[string]string{"msg": "file not found"}
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
