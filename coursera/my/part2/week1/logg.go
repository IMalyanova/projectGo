package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	addr := "localhost"
	port := 8080

	//std
	fmt.Printf("STD starting server at %s:%d", addr, port)

	//std
	log.Printf("STD starting server at %s%d", addr, port)

	//zap
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Info("starting server",
		zap.String("logger", "ZAP"),
		zap.String("host", addr),
		zap.Int("port", port),
	)

	//logrus
	logrus.SetFormatter(&logrus.TextFormatter{DisableColors: true})

	logrus.WithFields(logrus.Fields {
		"logger": "LOGRUS",
		"host":   addr,
		"port":   port,
	}).Info("Starting server")

	AccessLogOut := new(AccessLogger)

	//std
	AccessLogOut.StdLogger = log.New(os.Stdout, "STD", log.LUTC|log.Lshortfile)
}
