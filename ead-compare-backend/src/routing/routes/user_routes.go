package routes

import (
	"net/http"

	"github.com/m-rcussilva/go-private/tree/main/2024/projects/02-ead-compare/ead-compare-backend/src/controllers"
	"github.com/m-rcussilva/go-private/tree/main/2024/projects/02-ead-compare/ead-compare-backend/src/middleware"
)

var userRoutes = []RouteDefinitions{
	{
		URI:                   "/search",
		Method:                http.MethodPost,
		HandlerFunc:           controllers.SearchForCourses,
		RequireAuthentication: false,
	},
	{
		URI:                   "/register",
		Method:                http.MethodPost,
		HandlerFunc:           controllers.RegisterCourseUni,
		RequireAuthentication: true,
	},
	{
		URI:                   "/edit/{courseID}",
		Method:                http.MethodPut,
		HandlerFunc:           controllers.EditCourseUni,
		RequireAuthentication: true,
	},
	{
		URI:                   "/delete/{courseID}",
		Method:                http.MethodDelete,
		HandlerFunc:           controllers.DeleteCourseUni,
		RequireAuthentication: true,
	},
	{
		URI:                   "/list-all",
		Method:                http.MethodGet,
		HandlerFunc:           controllers.ListAll,
		RequireAuthentication: true,
	},
}

func ApplyAuthenticationMiddleware() {
	for i := range userRoutes {
		if userRoutes[i].RequireAuthentication {
			userRoutes[i].HandlerFunc = middleware.AdminAuthenticationMiddleware(http.HandlerFunc(userRoutes[i].HandlerFunc))
		}
	}
}
