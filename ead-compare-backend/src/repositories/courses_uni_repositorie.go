package repositories

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/m-rcussilva/go-private/tree/main/2024/projects/02-ead-compare/ead-compare-backend/src/models"
)

type CoursesUniRepo struct {
	db *sql.DB
}

func NewCourseUniRepository(db *sql.DB) *CoursesUniRepo {
	return &CoursesUniRepo{db}
}

func (c CoursesUniRepo) CreateCourse(course models.CursoUniversidade, disciplines []models.Disciplina) (uint64, error) {
	tx, err := c.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	result, err := tx.Exec(
		"INSERT INTO curso_universidade (nome_universidade, nota_mec, nome_curso, duracao, carga_horaria, formacao, informacoes_preco, link) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		course.NomeUniversidade,
		course.NotaMEC,
		course.NomeCurso,
		course.Duracao,
		course.CargaHoraria,
		course.Formacao,
		course.InformacoesPreco,
		course.Link,
	)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	fmt.Println("ID do curso inserido:", lastID)

	var courseID uint64
	err = tx.QueryRow("SELECT LAST_INSERT_ID()").Scan(&courseID)
	if err != nil {
		log.Println("Erro ao recuperar o último ID inserido:", err)
		return 0, err
	}

	for _, discipline := range disciplines {
		_, err := tx.Exec(
			"INSERT INTO disciplina (nome_disciplina, semestre) VALUES (?, ?)",
			discipline.NomeDisciplina, discipline.Semestre,
		)
		if err != nil {
			return 0, err
		}

		var disciplineID int
		err = tx.QueryRow("SELECT LAST_INSERT_ID()").Scan(&disciplineID)
		if err != nil {
			return 0, err
		}

		fmt.Println("ID da disciplina inserida:", disciplineID)

		_, err = tx.Exec(
			"INSERT INTO curso_disciplina (curso_id, disciplina_id, semestre) VALUES (?, ?, ?)",
			lastID, disciplineID, discipline.Semestre,
		)
		if err != nil {
			return 0, err
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return uint64(lastID), nil
}

func (c CoursesUniRepo) GetCourseName(courseName string) ([]models.CursoUniversidade, error) {
	courseName = fmt.Sprintf("%%%s%%", courseName)

	query := `
        SELECT cu.id, cu.nome_universidade, cu.nota_mec, cu.nome_curso, cu.duracao, cu.carga_horaria, cu.formacao, cu.informacoes_preco, cu.link,
               d.id, d.nome_disciplina, cd.semestre
        FROM curso_universidade cu
        JOIN curso_disciplina cd ON cu.id = cd.curso_id
        JOIN disciplina d ON cd.disciplina_id = d.id
        WHERE cu.nome_curso LIKE ?
    `

	rows, err := c.db.Query(query, courseName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []models.CursoUniversidade
	var currentCourse models.CursoUniversidade
	var currentDisciplina models.Disciplina

	for rows.Next() {
		err := rows.Scan(
			&currentCourse.ID,
			&currentCourse.NomeUniversidade,
			&currentCourse.NotaMEC,
			&currentCourse.NomeCurso,
			&currentCourse.Duracao,
			&currentCourse.CargaHoraria,
			&currentCourse.Formacao,
			&currentCourse.InformacoesPreco,
			&currentCourse.Link,
			&currentDisciplina.ID,
			&currentDisciplina.NomeDisciplina,
			&currentDisciplina.Semestre,
		)
		if err != nil {
			return nil, err
		}

		found := false
		for i := range courses {
			if courses[i].ID == currentCourse.ID {
				courses[i].Disciplinas = append(courses[i].Disciplinas, currentDisciplina)
				found = true
				break
			}
		}
		if !found {
			currentCourse.Disciplinas = []models.Disciplina{currentDisciplina}
			courses = append(courses, currentCourse)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}

func (c CoursesUniRepo) UpdateCourse(ID uint64, course models.CursoUniversidade) error {
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(
		"UPDATE curso_universidade SET nome_universidade = ?, nota_mec = ?, nome_curso = ?, duracao = ?, carga_horaria = ?, formacao = ?, informacoes_preco = ?, link = ? WHERE id = ?",
		course.NomeUniversidade, course.NotaMEC, course.NomeCurso, course.Duracao, course.CargaHoraria, course.Formacao, course.InformacoesPreco, course.Link, ID,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM curso_disciplina WHERE curso_id = ?", ID)
	if err != nil {
		return err
	}

	for _, discipline := range course.Disciplinas {
		_, err := tx.Exec(
			"INSERT INTO disciplina (nome_disciplina, semestre) VALUES (?, ?)",
			discipline.NomeDisciplina, discipline.Semestre,
		)
		if err != nil {
			return err
		}

		var disciplineID int
		err = tx.QueryRow("SELECT LAST_INSERT_ID()").Scan(&disciplineID)
		if err != nil {
			return err
		}

		_, err = tx.Exec("INSERT INTO curso_disciplina (curso_id, disciplina_id, semestre) VALUES (?, ?, ?)", ID, disciplineID, discipline.Semestre)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (c CoursesUniRepo) DeleteCourseUni(ID uint64) error {
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM curso_disciplina WHERE curso_id = ?", ID)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM curso_universidade WHERE id = ?", ID)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (c CoursesUniRepo) GetAllCourses() ([]models.CursoUniversidade, error) {
	query := `
        SELECT
            cu.id,
            cu.nome_universidade,
            cu.nota_mec,
            cu.nome_curso,
            cu.duracao,
            cu.carga_horaria,
            cu.formacao,
            cu.informacoes_preco,
            cu.link,
            d.id as disciplina_id,
            d.nome_disciplina,
            cd.semestre
        FROM curso_universidade cu
        LEFT JOIN curso_disciplina cd ON cu.id = cd.curso_id
        LEFT JOIN disciplina d ON cd.disciplina_id = d.id
    `
	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	coursesMap := make(map[uint64]*models.CursoUniversidade)
	for rows.Next() {
		var courseID uint64
		var course models.CursoUniversidade
		var discipline models.Disciplina
		err := rows.Scan(
			&courseID,
			&course.NomeUniversidade,
			&course.NotaMEC,
			&course.NomeCurso,
			&course.Duracao,
			&course.CargaHoraria,
			&course.Formacao,
			&course.InformacoesPreco,
			&course.Link,
			&discipline.ID,
			&discipline.NomeDisciplina,
			&discipline.Semestre,
		)
		if err != nil {
			return nil, err
		}
		if _, ok := coursesMap[courseID]; !ok {
			coursesMap[courseID] = &course
		}
		if discipline.ID != 0 {
			coursesMap[courseID].Disciplinas = append(coursesMap[courseID].Disciplinas, discipline)
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	var courses []models.CursoUniversidade
	for _, course := range coursesMap {
		courses = append(courses, *course)
	}
	return courses, nil
}
