package routes

import (
	"cli-project/internal/server/handlers"
	"github.com/gorilla/mux"
)

func InitialiseAuthRouter(r *mux.Router, authHandler *handlers.AuthHandler) {
	authRouter := r.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/signup", authHandler.SignupHandler).Methods("POST")
	authRouter.HandleFunc("/user/login", authHandler.UserLoginHandler).Methods("POST")
	authRouter.HandleFunc("/admin/login", authHandler.AdminLoginHandler).Methods("POST")
	authRouter.HandleFunc("/logout", authHandler.LogoutHandler).Methods("POST")
}
