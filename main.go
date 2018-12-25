package main

import (
	rawlog "log"
	"os"
	"logwrapper/logwrapper"
	gsrlog "github.com/go-forward/GSR/gsr/log"
)

func main() {
	rawLogger := rawlog.New(os.Stdout, "", rawlog.Lshortfile|rawlog.LstdFlags)
	logger := &logwrapper.Logger{}
	logger.SetRawLogger(rawLogger)

	var loggerImp gsrlog.LoggerInterface
	loggerImp = logger
	loggerImp.Errorf("Hello")
	loggerImp.Warningf("World")
}
