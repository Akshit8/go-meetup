package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Akshit8/go-meetup/db"
	"github.com/Akshit8/go-meetup/graph/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

type contextKey string

var currentUserKey = contextKey("currentUser")

// AuthMiddleware parses jwt token in request header and authorize it.
func AuthMiddleware(repo *db.UserRepo) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			token, err := parseToken(req)
			if err != nil {
				next.ServeHTTP(w, req)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)

			if !ok || !token.Valid {
				next.ServeHTTP(w, req)
				return
			}

			user, err := repo.GetUserByID(claims["jti"].(string))
			if err != nil {
				next.ServeHTTP(w, req)
				return
			}

			ctx := context.WithValue(req.Context(), currentUserKey, user)

			next.ServeHTTP(w, req.WithContext(ctx))
		})
	}
}

func stripBearerPrefixToken(token string) (string, error) {
	bearer := "BEARER"

	if len(token) > len(bearer) && strings.ToUpper(token[0:len(bearer)]) == bearer {
		return token[len(bearer)+1:], nil
	}

	return token, errors.New("error parsing token")
}

var authHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Authorization"},
	Filter:    stripBearerPrefixToken,
}

var authExtractor = &request.MultiExtractor{
	authHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

func parseToken(r *http.Request) (*jwt.Token, error) {
	jwtToken, err := request.ParseFromRequest(r, authExtractor, func(token *jwt.Token) (interface{}, error) {
		t := []byte(os.Getenv("JWT_SECRET"))
		return t, nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}
	return jwtToken, nil
}

// GetCurrentUserFromCTX return user attached to ctx if present.
func GetCurrentUserFromCTX(ctx context.Context) (*model.User, error) {
	errNoUserInContext := errors.New("no user in context")

	if ctx.Value(currentUserKey) == nil {
		return nil, errNoUserInContext
	}

	user, ok := ctx.Value(currentUserKey).(*model.User)
	if !ok || user.ID == "" {
		return nil, errNoUserInContext
	}

	return user, nil
}
