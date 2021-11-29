package graph

import (
	"context"
	"net/http"
	"time"

	"github.com/go-pg/pg/v9"
	"github.com/phifertoo/go_graphql_server/models"
)

const userloaderKey = "userloader"

func DataloaderMiddleware(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userloader := UserLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(ids []string) ([]*models.User, []error) {
				var users []*models.User
				// pg.In comes from the postgres
				err := db.Model(&users).Where("id in (?)", pg.In(ids)).Select()
				return users, []error{err}
			},
		}

		ctx := context.WithValue(r.Context(), userloaderKey, &userloader)

		// continue the next function
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserLoader(ctx context.Context) *UserLoader {
	return ctx.Value(userloaderKey).(*UserLoader)
}
