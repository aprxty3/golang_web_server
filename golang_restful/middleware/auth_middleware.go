package middleware

import (
	"golang_restful/helper"
	"golang_restful/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (authMiddleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	if "Rahasia" == request.Header.Get("X-API-Key") {

		authMiddleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webReponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized"}

		helper.WriteToResponseBody(writer, webReponse)
	}
}
