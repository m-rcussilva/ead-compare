package controllers

import (
	"encoding/json"
	"frontend/src/models"
	"frontend/src/utils"
	"log"
	"net/http"
)

func LoadHomePage(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "home.html", nil)
}

func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "login.html", nil)
}

func LoadRegisterPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "register.html", nil)
}

func LoadAboutPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "about.html", nil)
}

func LoadListOfAllCoursesPage(w http.ResponseWriter, r *http.Request) {
	log.Println("Recebida solicitação para carregar a página list-all.html")

	resp, err := http.Get("http://localhost:5000/list-all")
	if err != nil {
		log.Println("Erro ao fazer a solicitação para o backend:", err)
		http.Error(w, "Erro ao fazer a solicitação para o backend", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Erro ao recuperar os cursos do backend. Status code:", resp.StatusCode)
		http.Error(w, "Erro ao recuperar os cursos do backend", resp.StatusCode)
		return
	}

	var courses []models.Course
	if err := json.NewDecoder(resp.Body).Decode(&courses); err != nil {
		log.Println("Erro ao decodificar os dados dos cursos:", err)
		http.Error(w, "Erro ao decodificar os dados dos cursos", http.StatusInternalServerError)
		return
	}

	log.Println("Cursos recuperados com sucesso:", courses)

	utils.ExecTemplate(w, "list-all.html", nil)
}
