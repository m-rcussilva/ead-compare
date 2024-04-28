package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"frontend/src/responses"
	"io"
	"net/http"
)

func LogIn(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	admin, err := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		responses.JSONResponse(w, http.StatusBadRequest, responses.StructErrorForStatusCode{Err: err.Error()})
	}

	response, err := http.Post("http://localhost:5000/login", "application/json", bytes.NewBuffer(admin))
	if err != nil {
		responses.JSONResponse(w, http.StatusInternalServerError, responses.StructErrorForStatusCode{Err: err.Error()})
	}

	adminToken, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode, string(adminToken))
}
