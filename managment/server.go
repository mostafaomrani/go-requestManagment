package managment

import (
	"net/http"
	"time"

	"requestmanagment/handler"
)

func Run() {
	mux := http.NewServeMux()
	mux.Handle("/users/", &handler.UsersHandler{})

	server := http.Server{
		Addr:              ":8090",
		ReadHeaderTimeout: time.Second * 10,
		WriteTimeout:      time.Second * 10,
		Handler:           mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
