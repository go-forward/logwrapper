package main

import (
	"log"
	"os"
	"logwrapper/logwrapper"
	"github.com/go-forward/GSR/gsr"
)

func main() {
	rawLogger := log.New(os.Stdout, "", log.Lshortfile|log.LstdFlags)
	logger := &logwrapper.Logger{}
	logger.SetRawLogger(rawLogger)

	var loggerImp gsr.LoggerInterface
	loggerImp = logger
	loggerImp.Errorf("Hello")
	loggerImp.Warningf("World")
}
