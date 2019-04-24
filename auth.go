package gosaas

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dstpierre/gosaas/cache"
	"github.com/dstpierre/gosaas/data"
	"github.com/dstpierre/gosaas/model"
)

// Auth represents an authenticated user.
type Auth struct {
	AccountID int64
	UserID    int64
	Email     string
	Role      model.Roles
}

// Authenticator middleware used to authenticate requests.
//
// There are 4 ways to authenticate a request:
// 1. Via an HTTP header named X-API-KEY.
// 2. Via a querystring parameter named "key=token".
// 3. Via a cookie named X-API-KEY.
// 4. Via basic authentication.
//
// For routes with MinimumRole set as model.RolePublic there's no authentication performed.
func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("AUTHENTICATION MIDDLEWARE")

		ctx := r.Context()
		mr := ctx.Value(ContextMinimumRole).(model.Roles)

		key, pat, err := extractKeyFromRequest(r)
		fmt.Println("AUTH KEY", key)
		// if there's no authentication or an error
		if mr > model.RolePublic {
			if len(key) == 0 || err != nil {
				http.Redirect(w, r, "/users/login", http.StatusSeeOther)
				return
			}
		}

		ca := &cache.Auth{}

		// do we have this key on cache already
		var a Auth
		if err := ca.Exists(key, &a); err != nil {
			log.Println("error while trying to get cache auth", err)
		}

		if len(a.Email) > 0 {
			ctx = context.WithValue(ctx, ContextAuth, a)
		} else {
			// if the route required public access we do not
			// perform any authentication.
			if mr == model.RolePublic {
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			db := ctx.Value(ContextDatabase).(*data.DB)

			id, t := model.ParseToken(key)
			acct, usr, err := db.Users.Auth(id, t, pat)
			if err != nil {
				er := fmt.Sprintf("invalid token key: %v", err)
				http.Error(w, er, http.StatusUnauthorized)
				return
			}

			a.AccountID = acct.ID
			a.Email = usr.Email
			a.UserID = usr.ID
			a.Role = usr.Role

			// save it to cache
			ca.Set(key, a, 30*time.Minute)

			ctx = context.WithValue(ctx, ContextAuth, a)
		}

		// we authorize the request
		if a.Role < mr {
			http.Redirect(w, r, "/users/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func extractKeyFromRequest(r *http.Request) (key string, pat bool, err error) {
	// first let's look if the X-API-KEY is present in the HTTP header
	key = r.Header.Get("X-API-KEY")
	if len(key) > 0 {
		return
	}

	// check the query string
	key = r.URL.Query().Get("key")
	if len(key) > 0 {
		return
	}

	// check for cookie
	ck, er := r.Cookie("X-API-KEY")
	if er != nil {
		// If it's ErrNoCookie we must continue
		// otherwise this is a legit error
		if er != http.ErrNoCookie {
			err = er
			return
		}
	} else {
		key = ck.Value
		return
	}

	// check if we are supplying basic auth
	authorization := r.Header.Get("Authorization")
	s := strings.SplitN(authorization, " ", 2)
	if len(s) != 2 {
		err = fmt.Errorf("invalid basic authentication format: %s - you must provide Basic base64token", authorization)
		return
	}

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		err = fmt.Errorf("invalid basic authentication format: %s - you must provide Basic base64token", authorization)
		return
	}

	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		err = fmt.Errorf("invalid basic authentication, your token should be _:access_token - got %s", string(b))
		return
	}

	key = pair[1]
	pat = true

	return
}
