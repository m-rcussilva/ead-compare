package models

import (
	"errors"
	"strings"
)

type CursoUniversidade struct {
	ID               uint64       `json:"id"`
	NomeUniversidade string       `json:"nome_universidade"`
	NotaMEC          int          `json:"nota_mec"`
	NomeCurso        string       `json:"nome_curso"`
	Duracao          string       `json:"duracao"`
	CargaHoraria     string       `json:"carga_horaria"`
	Formacao         string       `json:"formacao"`
	InformacoesPreco string       `json:"informacoes_preco"`
	Link             string       `json:"link"`
	Disciplinas      []Disciplina `json:"disciplinas"`
}

func (course *CursoUniversidade) PrepareAndValidateCourse() error {
	course.TrimWhitesapceCourseFields()

	err := course.ValidateCourseFields()
	if err != nil {
		return err
	}

	return nil
}

func (course *CursoUniversidade) ValidateCourseFields() error {
	if course.NomeUniversidade == "" {
		return errors.New("o campo 'Nome da Universidade' é obrigatório")
	}

	if course.NotaMEC == 0 {
		return errors.New("o campo 'Nota MEC' é obrigatório")
	}

	if course.NomeCurso == "" {
		return errors.New("o campo 'Nome do curso' é obrigatório")
	}

	if course.Duracao == "" {
		return errors.New("o campo 'Duração' é obrigatório")
	}

	if course.CargaHoraria == "" {
		return errors.New("o campo 'Carga horária' é obrigatório")
	}

	if course.Formacao == "" {
		return errors.New("o campo 'Formação' é obrigatório")
	}

	if course.InformacoesPreco == "" {
		return errors.New("o campo 'Infrmação sobre o preço' é obrigatório")
	}

	if course.Link == "" {
		return errors.New("o campo 'Link' é obrigatório")
	}

	for _, disciplina := range course.Disciplinas {
		if disciplina.NomeDisciplina == "" {
			return errors.New("o campo 'Nome da Disciplina' é obrigatório")
		}
	}

	return nil
}

func (course *CursoUniversidade) TrimWhitesapceCourseFields() {
	course.NomeUniversidade = strings.TrimSpace(course.NomeUniversidade)
	course.NomeCurso = strings.TrimSpace(course.NomeCurso)
	course.Duracao = strings.TrimSpace(course.Duracao)
	course.CargaHoraria = strings.TrimSpace(course.CargaHoraria)
	course.Formacao = strings.TrimSpace(course.Formacao)
	course.InformacoesPreco = strings.TrimSpace(course.InformacoesPreco)
	course.Link = strings.TrimSpace(course.Link)
}
