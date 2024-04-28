package models

type CursoDisciplina struct {
	ID           uint64 `json:"id"`
	CursoID      uint64 `json:"curso_id"`
	DisciplinaID uint64 `json:"disciplina_id"`
}
