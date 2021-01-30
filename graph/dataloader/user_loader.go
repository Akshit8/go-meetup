package dataloader

import (
	"context"
	"log"
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
				
				log.Print("ids:", ids)
		
				err := db.Model(&users).Where("id in (?)", pg.In(ids)).Select()
				if err != nil {
					return nil, []error{err}
				}
		
				u := make(map[string]*model.User, len(users))
	
				for _, user := range users {
					u[user.ID] = user
				}

				result := make([]*model.User, len(ids))

				for i, id := range ids {
					result[i] = u[id]
				}

				return result, nil
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
