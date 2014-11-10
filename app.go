package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type data []map[string]string

func returnJSON(w http.ResponseWriter, r *http.Request) {
	// url := strings.Replace(r.URL.String(), "/", "", 1)

	log.Println("url " + r.URL.String())
	jsonFile, err := ioutil.ReadFile("./" + r.URL.String() + ".json")

	var p data

	_ = json.Unmarshal(jsonFile, &p)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(404) // HTTP 404
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

func main() {
	http.HandleFunc("/", returnJSON)
}
