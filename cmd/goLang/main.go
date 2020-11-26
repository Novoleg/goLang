package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	log := logrus.New()
	log.SetOutput(os.Stdout)

	log.Info("Старт приложения")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Fatal("Порт не задан")
	}
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	serv := http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: router,
	}
	go serv.ListenAndServe()
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	<-interrupt
	log.Info("Остановка приложения")

	timeout, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	err := serv.Shutdown(timeout)
	if err != nil {
		log.Error("Ошибка когда приложение упало: %v", err)
	}

	log.Info("Приложение остановлено")
}
