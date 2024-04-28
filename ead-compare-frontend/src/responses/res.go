package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

type StructErrorForStatusCode struct {
	Err string `json:"err"`
}

func JSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "/application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

// Trata as requisicoes com Status Code error 400 ou superior
func ProcessErrorStatusCode(w http.ResponseWriter, r *http.Response) {
	var err StructErrorForStatusCode
	json.NewDecoder(r.Body).Decode(&err)
	JSONResponse(w, r.StatusCode, err)
}
