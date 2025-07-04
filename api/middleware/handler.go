package middleware

import (
	"encoding/json"
	"github.com/tutul/book-server/service"
	"net/http"
	"time"
)

func GetTokenHandler(userService *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		usr, err := userService.Login(user, pass)
		if err != nil {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		exp := time.Now().Add(100 * time.Minute).Unix()

		_, token, _ := TokenAuth.Encode(map[string]interface{}{
			"user_id":  usr.Username, // Use username as user_id
			"username": usr.Username,
			"exp":      exp,
		})

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": token})
	}
}
