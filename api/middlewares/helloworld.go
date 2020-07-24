package middlewares

import (
	"github.com/tc-teams/fakefinder-crawler/api"
	"log"
	"net/http"
	"time"
)

//HelloWord is the first middlewares for api
func HelloWord() api.MiddlewareFunc {
	return func(h api.Handler) api.Handler {
		return func(w http.ResponseWriter, r *http.Request) *api.BaseError {
			log.Println("This request was sent at", time.Now())
		    return 	h(w, r)

		}
	}
}
