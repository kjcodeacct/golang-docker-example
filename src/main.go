package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/kjcodeacct/golang_docker_example/config"
	logger "github.com/kjcodeacct/golang_docker_example/logger"
	"go.uber.org/zap"
)

const processName = "goapp"

var zapLogger *zap.Logger

func setup() {
	var err error

	err = config.Setup(processName)
	if err != nil {
		log.Fatalln(err.Error())
	}

	logFileName := fmt.Sprintf("%s.log", processName)
	logFileName = filepath.Join(config.Get().LogDir, logFileName)
	err = logger.Setup(config.Get().LogMode, logFileName)
	if err != nil {
		log.Fatalln(err.Error())
	}

	zapLogger, err = logger.Get()
	if err != nil {
		log.Fatalln(err.Error())
	}

}

func main() {
	setup()

	log.Printf("starting application on port %d", config.Get().Port)
	http.HandleFunc("/hello-world", HelloWorldHandler)
	http.HandleFunc("/error-example", ErrorExampleHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Get().Port), nil)
}

// HelloWorldHandler
func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	HttpReturn(w, r, http.StatusOK, textContentType, "Hello World!")
}

// HelloWorldHandler
func ErrorExampleHandler(w http.ResponseWriter, r *http.Request) {
	err := errors.New("oh no, the gophers escaped!")
	HttpReturnErr(w, r, err)
}
