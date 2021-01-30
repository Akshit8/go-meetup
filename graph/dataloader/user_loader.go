package dataloader

import (
	"context"
	"net/http"
	"time"

	"github.com/Akshit8/go-meetup/graph/model"
	"github.com/go-pg/pg/v10"
)

const userLoaderKey = "userloader"

// Middleware loads users from db to optimise N + 1 queries
func Middleware(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userLoader := UserLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(ids []string) ([]*model.User, []error) {
				var users []*model.User
				
				err := db.Model(&users).Where("id in (?)", pg.In(ids)).Select()

				return users, []error{err}
			},
		}
		ctx := context.WithValue(r.Context(), userLoaderKey, &userLoader)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserLoader takes context and executes wrapped
func GetUserLoader(ctx context.Context) *UserLoader {
	return ctx.Value(userLoaderKey).(*UserLoader)
}
