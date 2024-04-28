package models

type Course struct {
	ID               uint64       `json:"id"`
	NomeUniversidade string       `json:"nome_universidade"`
	NotaMEC          int          `json:"nota_mec"`
	NomeCurso        string       `json:"nome_curso"`
	Duracao          string       `json:"duracao"`
	CargaHoraria     string       `json:"carga_horaria"`
	Formacao         string       `json:"formacao"`
	InformacoesPreco string       `json:"informacoes_preco"`
	Link             string       `json:"link"`
	Disciplinas      []Discipline `json:"disciplinas"`
}

type Discipline struct {
	ID             uint64 `json:"id"`
	NomeDisciplina string `json:"nome_disciplina"`
	Semestre       int    `json:"semestre"`
}

type CourseDiscipline struct {
	ID           uint64 `json:"id"`
	CourseID     uint64 `json:"curso_id"`
	DisciplineID uint64 `json:"disciplina_id"`
}
