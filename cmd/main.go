package main

import "github.com/sirupsen/logrus"

func main() {

	logrus.SetFormatter(&logrus.JSONFormatter{})

	//TODO init config

	//TODO init router

	//TODO run server
}
