package middleware

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
)

var (
	TokenAuth = jwtauth.New("HS256", []byte("your-secret-key"), nil)
	AdminUser = "AdminUser"
	AdminPass = "AdminPassword"
)

func GetTokenHandler(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()
	if !ok || user != AdminUser || pass != AdminPass {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	exp := time.Now().Add(100 * time.Minute).Unix()

	_, token, _ := TokenAuth.Encode(map[string]interface{}{
		"user_id":  108,
		"username": AdminUser,
		"exp":      exp,
	})

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
