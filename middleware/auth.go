package middleware

import (
	"net/http"

	"github.com/Akshit8/go-meetup/db"
)

const currentUserKey = "currentUser"

func AuthMiddleware(repo *db.UserRepo) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// TODO: token parsing and auth
		})
	}
}
