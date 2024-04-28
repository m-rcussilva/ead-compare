package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/m-rcussilva/go-private/tree/main/2024/projects/02-ead-compare/ead-compare-backend/src/database"
	"github.com/m-rcussilva/go-private/tree/main/2024/projects/02-ead-compare/ead-compare-backend/src/models"
	"github.com/m-rcussilva/go-private/tree/main/2024/projects/02-ead-compare/ead-compare-backend/src/repositories"
	"github.com/m-rcussilva/go-private/tree/main/2024/projects/02-ead-compare/ead-compare-backend/src/responses"
)

func SearchForCourses(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		CourseName string `json:"course"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		log.Printf("Erro ao decodificar solicitação: %v", err)
		responses.JSONErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	courseName := strings.ToLower(requestData.CourseName) // Linha adicionada

	db, err := database.ConnectionDB()
	if err != nil {
		log.Printf("Erro ao conectar ao banco de dados: %v", err)
		responses.JSONErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorie := repositories.NewCourseUniRepository(db)
	courses, err := repositorie.GetCourseName(courseName)
	if err != nil {
		log.Printf("Erro ao buscar cursos no banco de dados: %v", err)
		responses.JSONErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSONResponse(w, http.StatusOK, courses)
}

func RegisterCourseUni(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var requestData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		log.Println("Erro ao decodificar JSON do curso:", err)
		responses.JSONErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	newCourse := models.CursoUniversidade{
		NomeUniversidade: requestData["nome_universidade"].(string),
		NomeCurso:        requestData["nome_curso"].(string),
		Duracao:          requestData["duracao"].(string),
		CargaHoraria:     requestData["carga_horaria"].(string),
		Formacao:         requestData["formacao"].(string),
		InformacoesPreco: requestData["informacoes_preco"].(string),
		Link:             requestData["link"].(string),
	}

	var notaMEC string
	if val, ok := requestData["nota_mec"].(float64); ok {
		notaMEC = fmt.Sprintf("%.0f", val)
	} else if valStr, ok := requestData["nota_mec"].(string); ok {
		notaMEC = valStr
	} else {
		log.Println("Erro: o campo 'Nota MEC' é inválido")
		responses.JSONErrorResponse(w, http.StatusBadRequest, errors.New("o campo 'Nota MEC' é inválido"))
		return
	}
	var err error
	newCourse.NotaMEC, err = strconv.Atoi(notaMEC)
	if err != nil {
		log.Println("Erro ao converter 'Nota MEC' para inteiro:", err)
		responses.JSONErrorResponse(w, http.StatusBadRequest, errors.New("o campo 'Nota MEC' é inválido"))
		return
	}

	if newCourse.NomeUniversidade == "" {
		log.Println("Erro: o campo 'Nome da Universidade' é obrigatório")
		responses.JSONErrorResponse(w, http.StatusBadRequest, errors.New("o campo 'Nome da Universidade' é obrigatório"))
		return
	}

	disciplinasData := requestData["disciplinas"].([]interface{})
	var disciplinas []models.Disciplina
	for _, disc := range disciplinasData {
		discMap := disc.(map[string]interface{})
		nomeDisciplinas := toStringSlice(discMap["nome_disciplina"].([]interface{}))
		disciplinas = append(disciplinas, models.Disciplina{
			NomeDisciplina: strings.Join(nomeDisciplinas, ", "),
			Semestre:       int(discMap["semestre"].(float64)),
		})
	}

	newCourse.Disciplinas = disciplinas
	log.Println("Curso validado com sucesso:", newCourse)

	if err := newCourse.PrepareAndValidateCourse(); err != nil {
		log.Println("Erro ao validar curso:", err)
		responses.JSONErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.ConnectionDB()
	if err != nil {
		log.Println("Erro ao conectar ao banco de dados:", err)
		responses.JSONErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewCourseUniRepository(db)
	courseID, err := repository.CreateCourse(newCourse, newCourse.Disciplinas)
	if err != nil {
		log.Println("Erro ao criar curso no banco de dados:", err)
		responses.JSONErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	successResponse := responses.SuccessResponse{
		Message: "Curso criado com sucesso",
		ID:      int(courseID),
	}

	log.Println("Curso criado com sucesso:", successResponse)

	responses.JSONResponse(w, http.StatusCreated, successResponse)
}

// Função auxiliar para converter []interface{} em []string
func toStringSlice(interfaces []interface{}) []string {
	strings := make([]string, len(interfaces))
	for i, v := range interfaces {
		strings[i] = fmt.Sprintf("%v", v)
	}
	return strings
}

func EditCourseUni(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	courseID, err := strconv.ParseUint(params["courseID"], 10, 64)
	if err != nil {
		responses.JSONErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.JSONErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	var courseUni models.CursoUniversidade

	if err = json.Unmarshal(bodyRequest, &courseUni); err != nil {
		responses.JSONErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if len(courseUni.Disciplinas) == 0 {
		responses.JSONErrorResponse(w, http.StatusBadRequest, errors.New("o curso deve conter pelo menos uma disciplina"))
		return
	}

	if err = courseUni.PrepareAndValidateCourse(); err != nil {
		responses.JSONErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.ConnectionDB()
	if err != nil {
		responses.JSONErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewCourseUniRepository(db)
	if err = repository.UpdateCourse(courseID, courseUni); err != nil {
		responses.JSONErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSONResponse(w, http.StatusNoContent, nil)
}

func DeleteCourseUni(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	courseID, err := strconv.ParseUint(params["courseID"], 10, 64)
	if err != nil {
		responses.JSONErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.ConnectionDB()
	if err != nil {
		responses.JSONErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorie := repositories.NewCourseUniRepository(db)
	if err = repositorie.DeleteCourseUni(courseID); err != nil {
		responses.JSONErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSONResponse(w, http.StatusNoContent, nil)
}

func ListAll(w http.ResponseWriter, r *http.Request) {
	db, err := database.ConnectionDB()
	if err != nil {
		responses.JSONErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewCourseUniRepository(db)
	courses, err := repository.GetAllCourses()
	if err != nil {
		responses.JSONErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSONResponse(w, http.StatusOK, courses)
}
