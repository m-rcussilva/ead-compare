package routes

import (
	"frontend/src/controllers"
	"net/http"
)

var registeredRoutes = []Route{
	{
		URI:                   "/list-all-courses",
		Method:                http.MethodGet,
		HandlerFunc:           controllers.LoadListOfAllCoursesPage,
		RequireAuthentication: true,
	},
}
