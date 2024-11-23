package main

import "github.com/sirupsen/logrus"

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	//TODO init config
	//TODO init db
	//TODO init grpc services
	//TODO init repo
	//TODO init service
	//TODO init handler
	//TODO init router
	//TODO run server
	//TODO graceful shutdown
}
