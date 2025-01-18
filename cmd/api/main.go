package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/Tanyabatra541/GoLang/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	r := chi.NewRouter()
	handlers.Handler(r)
	fmt.Println("starting GO API service...")
	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}

