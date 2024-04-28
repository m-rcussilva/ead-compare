package routes

import (
	"net/http"

	"github.com/m-rcussilva/go-private/tree/main/2024/projects/02-ead-compare/ead-compare-backend/src/config"
)

var loginRoute = []RouteDefinitions{
	{
		URI:                   "/login",
		Method:                http.MethodPost,
		HandlerFunc:           config.HandleLogin,
		RequireAuthentication: false,
	},
}
