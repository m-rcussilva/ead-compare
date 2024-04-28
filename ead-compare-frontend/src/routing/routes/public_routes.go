package routes

import (
	"frontend/src/controllers"
	"net/http"
)

var publicRoutes = []Route{
	{
		URI:                   "/",
		Method:                http.MethodGet,
		HandlerFunc:           controllers.LoadHomePage,
		RequireAuthentication: false,
	},
	{
		URI:                   "/login",
		Method:                http.MethodGet,
		HandlerFunc:           controllers.LoadLoginPage,
		RequireAuthentication: false,
	},
	{
		URI:                   "/login",
		Method:                http.MethodPost,
		HandlerFunc:           controllers.LogIn,
		RequireAuthentication: false,
	},
	{
		URI:                   "/about",
		Method:                http.MethodGet,
		HandlerFunc:           controllers.LoadAboutPage,
		RequireAuthentication: false,
	},
}
