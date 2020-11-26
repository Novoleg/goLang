package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	logrus.Info("Всем привет, я только проснулся")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		logrus.Fatal("Порт не задан")
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	http.ListenAndServe(":"+port, nil)
}
