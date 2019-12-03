package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/leftpad", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Is GET correct here?
		if r.Method != "GET" {
			log.Printf("%s on %s is not allowed\n", r.Method, "/leftpad")
			http.Error(w, http.StatusText(405), 405)
			return
		}

		// Read without buffer
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalln(err)
		}

		var dat map[string]interface{}
		if err := json.Unmarshal(body, &dat); err != nil {
			respondWithJSON(w, http.StatusInternalServerError, "")
		}
		log.Println(dat)
		
		// respond
		respondWithJSON(w, http.StatusOK, dat["txt"])
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.MarshalIndent(payload, "", " ")
	if err != nil {
		log.Printf("%s could not be marshaled into JSON: %s", data, err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}