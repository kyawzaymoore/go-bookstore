package middlewares

import (
	"net/http"
)

var Users = map[string]string{
    "fido": "dido",
}

// BasicAuthMiddleware checks for Basic Authentication
func BasicAuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        username, password, ok := r.BasicAuth()
        if !ok || !checkUser(username, password) {
            w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r)
    })
}

// checkUser validates the username and password for Basic Auth
func checkUser(username, password string) bool {
    if pass, ok := Users[username]; ok {
        return pass == password
    }
    return false
}
