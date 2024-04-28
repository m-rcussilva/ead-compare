package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

type SuccessResponse struct {
	Message string `json:"message"`
	ID      int    `json:"id"`
}

func JSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

func JSONErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	JSONResponse(w, statusCode, struct {
		Err string `json:"err"`
	}{
		Err: err.Error(),
	})
}
