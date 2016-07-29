package log

import (
	"log"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *log.Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
var pro string
var process = false
var debug = false

func SetProcess(p string) {

	log.SetOutput(&lumberjack.Logger{
		Filename:   "/var/log/axihome/" + p + ".log",
		MaxSize:    5, // megabytes
		MaxBackups: 10,
		MaxAge:     28, //days
	})

	pro = p
	if pro != "" {

		process = true
	} else {

		process = false
	}
}

func SetDebug(b bool) {

	debug = b
}

func Println(a ...interface{}) {

	log.Println(a...)

	if process {
		logger.Printf("[%s] : %v", pro, a[0])
	} else {

		logger.Println(a...)
	}

}

func Debug(a ...interface{}) {

	if debug {

		log.Println(a...)

		if process {
			logger.Printf("[%s] : %v", pro, a[0])
		} else {

			logger.Println(a...)
		}
	}
}
