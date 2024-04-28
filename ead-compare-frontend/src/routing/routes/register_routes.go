package routes

import (
	"frontend/src/controllers"
	"net/http"
)

var registerRoutes = []Route{
	{
		URI:                   "/register-page",
		Method:                http.MethodGet,
		HandlerFunc:           controllers.LoadRegisterPage,
		RequireAuthentication: true,
	},
	{
		URI:                   "/register",
		Method:                http.MethodPost,
		HandlerFunc:           controllers.CreateCourse,
		RequireAuthentication: true,
	},
}
