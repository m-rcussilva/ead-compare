package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"frontend/src/responses"
	"log"
	"net/http"
)

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	log.Println("Recebida solicitação para criar um novo curso")

	r.ParseForm()

	var requestData map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		log.Println("Erro ao decodificar JSON do curso:", err)
		http.Error(w, "Erro ao decodificar JSON do curso", http.StatusBadRequest)
		return
	}

	log.Println("Dados do curso recebidos:", requestData)

	disciplinasJSON, err := json.Marshal(requestData["disciplinas"])
	if err != nil {
		log.Println("Erro ao serializar JSON de disciplinas:", err)
		http.Error(w, "Erro ao serializar JSON de disciplinas", http.StatusInternalServerError)
		return
	}

	var disciplinas []map[string]interface{}
	if err := json.Unmarshal(disciplinasJSON, &disciplinas); err != nil {
		log.Println("Erro ao decodificar JSON de disciplinas:", err)
		http.Error(w, "Erro ao decodificar JSON de disciplinas", http.StatusBadRequest)
		return
	}

	fmt.Println("Dados do curso:", requestData)
	fmt.Println("Disciplinas:", disciplinas)

	notaMECFloat64, ok := requestData["nota_mec"].(float64)
	if !ok {
		log.Println("Erro: Valor da nota MEC inválido")
		http.Error(w, "Erro: Valor da nota MEC inválido", http.StatusBadRequest)
		return
	}

	notaMEC := int(notaMECFloat64)

	course := map[string]interface{}{
		"nome_universidade": requestData["nome_universidade"].(string),
		"nota_mec":          notaMEC,
		"nome_curso":        requestData["nome_curso"].(string),
		"duracao":           requestData["duracao"].(string),
		"carga_horaria":     requestData["carga_horaria"].(string),
		"formacao":          requestData["formacao"].(string),
		"informacoes_preco": requestData["informacoes_preco"].(string),
		"link":              requestData["link"].(string),
		"disciplinas":       disciplinas,
	}

	courseJSON, err := json.Marshal(course)
	if err != nil {
		log.Println("Erro ao serializar dados do curso para JSON:", err)
		http.Error(w, "Erro ao serializar dados do curso para JSON", http.StatusInternalServerError)
		return
	}

	fmt.Println(string(courseJSON))

	response, err := http.Post("http://localhost:5000/register", "application/json", bytes.NewBuffer(courseJSON))
	if err != nil {
		responses.JSONResponse(w, http.StatusInternalServerError, responses.StructErrorForStatusCode{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.ProcessErrorStatusCode(w, response)
		return
	}

	responses.JSONResponse(w, response.StatusCode, nil)
}
