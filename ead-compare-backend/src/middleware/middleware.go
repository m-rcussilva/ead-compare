package middleware

import (
	"net/http"
)

func AdminAuthenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !isAdminAuthenticated(r) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Erro ao analisar o formulário", http.StatusBadRequest)
			return
		}

		next(w, r)
	}
}

func isAdminAuthenticated(r *http.Request) bool {
	_, err := r.Cookie("admin_authenticated")
	return err == nil
}
