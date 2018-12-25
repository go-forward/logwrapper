package main

import (
	rawlog "log"
	"os"
	"logwrapper/logwrapper"
	"github.com/go-forward/GSR/gsr"
)

func main() {
	rawLogger := rawlog.New(os.Stdout, "", rawlog.Lshortfile|rawlog.LstdFlags)
	logger := &logwrapper.Logger{}
	logger.SetRawLogger(rawLogger)

	var loggerImp gsr.LoggerInterface
	loggerImp = logger
	loggerImp.Errorf("Hello")
	loggerImp.Warningf("World")
}
