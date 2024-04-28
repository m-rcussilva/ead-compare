var typingTimer;
var doneTypingInterval = 500;
const endpointURL = "http://localhost:5000/search"

$(document).ready(function() {
    $("#course").on("input", function() {
        clearTimeout(typingTimer);
        if ($(this).val()) {
            typingTimer = setTimeout(searchForCourses, doneTypingInterval);
        }
    });

    $("#suggestions").on("click", "li", function() {
        var selectedCourse = $(this).text();
        $("#course").val(selectedCourse);
        $("#suggestions").empty();
    });
});

function searchForCourses() {
    var query = $("#course").val();
    $.ajax({
        url: endpointURL,
        method: "POST",
        contentType: "application/json",
        data: JSON.stringify({ course: query }),
        success: function(response) {
            $("#suggestions").empty();
            var coursesHTML = '';
            response.forEach(function(course, index) {
                var courseCard = `
                            <div class="course-card">
                                <h2>${course.nome_curso}</h2>
                                <p><strong>Universidade:</strong> ${course.nome_universidade}</p>
                                <p><strong>Nota MEC:</strong> ${course.nota_mec}</p>
                                <p><strong>Duração:</strong> ${course.duracao}</p>
                                <p><strong>Carga horária:</strong> ${course.carga_horaria}</p>
                                <p><strong>Formação:</strong> ${course.formacao}</p>
                                <p><strong>Informações preço:</strong> ${course.informacoes_preco}</p>
                                <p><strong>Link:</strong> ${course.link}</p>
                                <p><strong>Disciplinas:</strong></p>
                                <ul>
                        `;
                course.disciplinas.forEach(function(disciplina) {
                    courseCard += `<li>${disciplina.nome_disciplina} (Semestre ${disciplina.semestre})</li>`;
                });
                courseCard += `
                                </ul>
                            </div>
                        `;
                coursesHTML += courseCard;
                if ((index + 1) % 2 === 0) {
                    coursesHTML += '<div style="clear:both;"></div>';
                }
            });
            $("#suggestions").html(coursesHTML);
        },
        error: function(xhr, status, error) {
            console.error(xhr, status, error);
            $("#suggestions").empty();
        }
    });
}
