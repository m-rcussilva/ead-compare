package models

import (
	"errors"
	"strconv"
	"strings"
)

type Disciplina struct {
	ID             uint64 `json:"id"`
	NomeDisciplina string `json:"nome_disciplina"`
	Semestre       int    `json:"semestre"`
}

func (discipline *Disciplina) PrepareAndValidateDiscipline() error {
	discipline.TrimWhitespaceDisciplineField()

	err := discipline.ValdiateDisciplineField()
	if err != nil {
		return err
	}

	return nil
}

func (discipline *Disciplina) ValdiateDisciplineField() error {
	if discipline.NomeDisciplina == "" {
		return errors.New("o campo 'Nome das disciplinas' é obrigatório")
	}

	if strconv.Itoa(discipline.Semestre) == "" {
		return errors.New("o campo 'Semestre' é obrigatório")
	}

	return nil
}

func (discipline *Disciplina) TrimWhitespaceDisciplineField() {
	discipline.NomeDisciplina = strings.TrimSpace(discipline.NomeDisciplina)
}
