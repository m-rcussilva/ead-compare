const endpointURL = "http://localhost:5000/list-all";

$(document).ready(function() {
    console.log("Iniciando solicitação para listar todos os cursos")
    $.ajax({
        url: endpointURL,
        type: 'GET',
        dataType: 'json',
        success: function(data) {
            console.log('Dados dos cursos recebidos:', data);
            setTimeout(function() {
                renderCourses(data);
            }, 1000)
        },
        error: function(xhr, status, error) {
            console.error('Erro durante a solicitação dos cursos:', xhr, status, error);
            displayErrorMessage('Erro ao carregar os cursos. Por favor, tente novamente mais tarde.');
        }
    });
});

function renderCourses(courses) {
    if (courses.length === 0) {
        $('#course-list').append('<p>Nenhum curso encontrado.</p>');
        return;
    }

    courses.forEach(function(course) {
        console.log('Curso:', course);
        console.log("ID DO CURSO:", course.id)

        var courseHTML = '<div class="course-item">';
        courseHTML += '<h3>' + (course.nome_curso || 'Nome do curso não disponível') + '</h3>';
        courseHTML += '<p><strong>Universidade:</strong> ' + (course.nome_universidade || 'Não disponível') + '</p>';
        courseHTML += '<p><strong>Nota MEC:</strong> ' + (course.nota_mec || 'Não disponível') + '</p>';
        courseHTML += '<p><strong>Duração:</strong> ' + (course.duracao || 'Não disponível') + '</p>';
        courseHTML += '<p><strong>Carga Horária:</strong> ' + (course.carga_horaria || 'Não disponível') + '</p>';
        courseHTML += '<p><strong>Formação:</strong> ' + (course.formacao || 'Não disponível') + '</p>';
        courseHTML += '<p><strong>Informações preço:</strong> ' + (course.informacoes_preco || 'Não disponível') + '</p>';
        courseHTML += '<p><strong>Link:</strong> ' + (course.link || 'Não disponível') + '</p>';
        if (course.disciplinas && course.disciplinas.length > 0) {
            courseHTML += '<p><strong>Disciplinas:</strong></p>';
            courseHTML += '<ul>';
            course.disciplinas.forEach(function(disciplina) {
                console.log('Disciplina:', disciplina);
                courseHTML += '<li>' + (disciplina.nome_disciplina || 'Nome da disciplina não disponível') + ' (Semestre ' + (disciplina.semestre || 'Não especificado') + ')</li>';
            });
            courseHTML += '</ul>';
        }
        if (course.id) {
            console.log('Adicionando botões para o curso com ID:', course.id);
            courseHTML += '<button class="edit-button" onclick="editCourse(' + course.id + ')">Editar</button>';
            courseHTML += '<button class="delete-button" onclick="deleteCourse(' + course.id + ')">Excluir</button>';
        }
        courseHTML += '</div>';

        $('#course-list').append(courseHTML);
    });
}

function displayErrorMessage(message) {
    $('#error-message').text(message);
}

function editCourse(courseId) {
    console.log('Editando curso com ID:', courseId);
}

function deleteCourse(courseId) {
    console.log('Excluindo curso com ID:', courseId);
}
