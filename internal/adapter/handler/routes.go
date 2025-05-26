package handler

import "net/http"

func NewRouter(userHandler *UserHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			userHandler.Save(w, r)
			return
		}
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	})

	mux.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			userHandler.FindByID(w, r)
			return
		}
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	})

	return mux
}
