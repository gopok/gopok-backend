package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/gopok/gopok-backend/pkg/core"
	"gopkg.in/mgo.v2/bson"
)

type authContextKeyType string

/*
UserContextKey is used to retrieve the user from request context
*/
var UserContextKey = authContextKeyType("UserContextKey")

/*
SessionContextKey is used to retrieve the session from request context
*/
var SessionContextKey = authContextKeyType("SessionContextKey")

func extractSessionToken(r *http.Request) (string, error) {
	authData := r.Header.Get("Authorization")
	if len(authData) == 0 || !strings.HasPrefix(authData, "Bearer ") {

		return "", errors.New("Authorization header is empty, missing or has no 'Bearer ' prefix")
	}
	return strings.TrimPrefix(authData, "Bearer "), nil
}

/*
CheckUserMiddleware validates the session token passed in the Authorization header, aborts the request if it is invalid. When a session is found it attaches the user to the context.
*/
func CheckUserMiddleware(app *core.Application) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sessToken, extractionErr := extractSessionToken(r)
			if extractionErr != nil {
				w.WriteHeader(401)
				jsonErr, _ := json.Marshal(core.NewErrorResponse(extractionErr.Error(), 401))
				w.Write(jsonErr)
				return
			}
			session := &session{}
			findErr := app.Db.C("sessions").Find(bson.M{
				"token": sessToken,
			}).One(session)
			if findErr != nil {
				w.WriteHeader(401)
				jsonErr, _ := json.Marshal(core.NewErrorResponse("Failed to find session: "+findErr.Error(), 401))
				w.Write(jsonErr)
				return
			}
			if !session.Active || session.ExpiresOn.Before(time.Now()) {
				w.WriteHeader(401)
				jsonErr, _ := json.Marshal(core.NewErrorResponse("Session inactive or expired", 401))
				w.Write(jsonErr)
				return
			}
			user := &User{}
			userFindErr := app.Db.C("users").FindId(session.UserID).One(user)
			if userFindErr != nil {
				w.WriteHeader(401)
				jsonErr, _ := json.Marshal(core.NewErrorResponse("Failed to find user: "+userFindErr.Error(), 401))
				w.Write(jsonErr)
				return
			}
			ctx := r.Context()
			ctx = context.WithValue(ctx, SessionContextKey, session)
			ctx = context.WithValue(ctx, UserContextKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}

}
