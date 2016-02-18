package log

import (
	"log"
	"os"
)

var logger *log.Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
var pro string
var process = false
var debug = false

func SetProcess(p string) {

	//logger = log.New(os.Stdout, p+"\t", log.Ldate|log.Ltime)
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

	if process {
		logger.Printf("[%s] : %v", pro, a[0])
	} else {

		logger.Println(a...)
	}

}

func Debug(a ...interface{}) {

	if debug {
		if process {
			logger.Printf("[%s] : %v", pro, a[0])
		} else {

			logger.Println(a...)
		}
	}
}
