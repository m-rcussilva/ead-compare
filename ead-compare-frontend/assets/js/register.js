$("#form-register").on("submit", createCourseUni);

let disciplineCount = 1;
let disciplinas = [];

function addDiscipline() {
    disciplineCount++;
    const divDiscipline = document.createElement('div');
    divDiscipline.classList.add('disc-flex-column');
    divDiscipline.innerHTML = `<label for="disciplina_${disciplineCount}">Disciplina do ${disciplineCount}º Semestre:</label>
                                        <input type="text" id="discipline_${disciplineCount}" name="disciplinas[]" required>`;
    document.getElementById('disciplinas').appendChild(divDiscipline);
}

function createCourseUni(e) {
    e.preventDefault();

    const formData = new FormData(document.getElementById("form-register"));

    console.log("Formulário de registro enviado com sucesso. Dados do formulário:", formData)

    if (!formData.has('disciplinas[]')) {
        console.error("Erro: Nenhuma disciplina foi adicionada.");
        alert("Adicione pelo menos uma disciplina.");
        return;
    }

    const requestData = {
        nome_universidade: formData.get("nome_universidade"),
        nota_mec: parseInt(formData.get("nota_mec")),
        nome_curso: formData.get("nome_curso"),
        duracao: formData.get("duracao"),
        carga_horaria: formData.get("carga_horaria"),
        formacao: formData.get("formacao"),
        informacoes_preco: formData.get("informacoes_preco"),
        link: formData.get("link"),
        disciplinas: Array.from(formData.getAll('disciplinas[]')).map((disciplina, index) => ({
            id: index + 1,
            nome_disciplina: disciplina.split(',').map(str => str.trim()),
            semestre: index + 1
        })),
    };

    console.log(JSON.stringify(requestData));

    $.ajax({
        url: "/register",
        method: "POST",
        contentType: "application/json",
        data: JSON.stringify(requestData),
    }).done(function(response) {
        console.log("Resposta do servidor:", response);
        alert("Curso cadastrado com sucesso!");
    }).fail(function(xhr, status, error) {
        console.error("Erro ao cadastrar um novo curso:", xhr, status, error);
        alert("Erro ao cadastrar um novo curso.");
    });
}
