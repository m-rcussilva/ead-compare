$("#login-form").on("submit", login)

function login(e) {
    e.preventDefault()

    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: $("#email").val(),
            password: $("#password").val(),
        }
    }).done(function(response) {
        console.log("Sucesso:", response)
        window.location = "/register-page"
    }).fail(function(status, error) {
        console.error(status, error)
        alert("Credenciais inválidas.")
    })
}
