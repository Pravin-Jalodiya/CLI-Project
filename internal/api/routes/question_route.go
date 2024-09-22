package routes

import (
	"cli-project/internal/api/handlers"
	"cli-project/internal/api/middleware"
	"github.com/gorilla/mux"
)

func InitialiseQuestionRouter(r *mux.Router, questionHandler *handlers.QuestionHandler) {
	r.HandleFunc("/question", questionHandler.GetQuestions).Methods("GET")
	questionRouter := r.PathPrefix("/question").Subrouter()
	questionRouter.Use(middleware.JWTAuthMiddleware)
	questionRouter.Use(middleware.AdminRoleMiddleware)
	questionRouter.HandleFunc("", questionHandler.RemoveQuestionById).Methods("DELETE")
	questionRouter.HandleFunc("/add-questions", questionHandler.AddQuestions).Methods("POST")
}
