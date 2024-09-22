package middleware

import (
	"cli-project/internal/config/roles"
	"net/http"
)

func AdminRoleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userMetaData, ok := r.Context().Value("userMetaData").(UserMetaData)
		if !ok || userMetaData.Role != roles.ADMIN {
			unauthorized(w, "Unauthorized access")
			return
		}
		next.ServeHTTP(w, r)
	})
}
