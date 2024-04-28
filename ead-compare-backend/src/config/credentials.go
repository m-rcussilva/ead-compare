package config

import (
	"encoding/json"
	"log"
	"net/http"
)

var (
	AdminEmail    = "admin@email.com"
	AdminPassword = "123456"
)

type LoginResponse struct {
	Message           string `json:"message"`
	AllowRegisterPage bool   `json:"allow_register_page"`
}

type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

/*
func StartLoginServer(port string) {
	http.HandleFunc("/login", HandleLogin)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
*/

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var formData LoginForm
	if err := json.NewDecoder(r.Body).Decode(&formData); err != nil {
		http.Error(w, "Erro ao decodificar os dados do formulário", http.StatusInternalServerError)
		log.Printf("Erro ao decodificar os dados do formulário: %v", err)
		return
	}

	log.Println("Dados do formulário:", formData)

	email := formData.Email
	password := formData.Password

	log.Println("Tentativa de login com o Email:", email)

	if email == "" || password == "" {
		http.Error(w, "Email e senha são obrigatórios", http.StatusBadRequest)
		return
	}

	if email != AdminEmail || password != AdminPassword {
		log.Println("Credenciais inválidas para o email:", email)
		http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
		return
	}

	if email == AdminEmail && password == AdminPassword {
		log.Println("Login bem-sucedido para o email:", email)
		http.SetCookie(w, &http.Cookie{
			Name:   "admin_authenticated",
			Value:  "true",
			MaxAge: 3600,
		})

		response := LoginResponse{
			Message:           "Login efetuado com sucesso!",
			AllowRegisterPage: true,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	log.Println("Credenciais inválidas para o email:", email)
	http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
}
