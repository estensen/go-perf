package profiling

import (
	"encoding/json"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strconv"
)

func main() {
	leftpadHandler()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func leftpadHandler() {
	// TODO: Do something mega-slow that will dominate CPU time

	http.HandleFunc("/leftpad", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Incoming request: ", r.RequestURI)

		if r.Method != "GET" {
			log.Printf("%s on %s is not allowed\n", r.Method, "/leftpad")
			http.Error(w, http.StatusText(405), 405)
			return
		}

		str := r.FormValue("str")
		if str == "" {
			log.Println("no string")
			http.Error(w, http.StatusText(400), 400)
			return
		}

		padding := r.FormValue("length")
		if padding == "" {
			log.Println("no length")
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		str, err := leftpad(str, padding)
		if err != nil {
			log.Println("padding NaN")
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		respondWithJSON(w, http.StatusOK, str)
	})
}

func leftpad(str string, length string) (string, error) {
	l, err := strconv.Atoi(length)
	if err != nil {
		return "", err
	}

	for len(str) < l {
		str = " " + str
	}
	return str, nil
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
