package v1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
func jsonResponse(w http.ResponseWriter, data interface{}, c int) {
	//	dj, err := json.MarshalIndent(data, "", "  ")
	//	if err != nil {
	//	http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
	//	log.Println(err)
	//	return
	//}
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c)
	fmt.Fprintf(w, "%s", data)
}

func jsonResponseStruct(w http.ResponseWriter, data interface{}, c int) {
	dj, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c)
	fmt.Fprintf(w, "%s", dj )
}
