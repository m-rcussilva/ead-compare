package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/m-rcussilva/go-private/tree/main/2024/projects/02-ead-compare/ead-compare-backend/src/config"
	"github.com/m-rcussilva/go-private/tree/main/2024/projects/02-ead-compare/ead-compare-backend/src/responses"
)

func LogIn(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.JSONErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	var requestBody map[string]interface{}
	if err := json.Unmarshal(bodyRequest, &requestBody); err != nil {
		responses.JSONErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	config.HandleLogin(w, r)
}
