package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	logrus.Info("Всем привет, я только проснулся")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		logrus.Fatal("Порт не задан")
	}
}
