package main

import (
	"fmt"
	"github.com/I-Limar/link_shortening/app"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {

	var conf app.Config

	envconfig.MustProcess("", &conf)
	h, err := app.NewApp(&conf)
	checkError(err, "Init app")

	addr := fmt.Sprintf(":%d", conf.HTTPPort)
	logrus.Infof("🚀  The server is up and running on localhost:%d", conf.HTTPPort)

	checkError(http.ListenAndServe(addr, h), "the server listening return error")
}

func checkError(err error, msg string) {
	if err != nil {
		logrus.Error(msg, ":", err)
		os.Exit(1)
	}
}
