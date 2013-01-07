package kuler

import (
	"io"
	"log"
)

var logger *log.Logger

func LogTo(out io.Writer) {
	// io/ioutil.Discard
	// os.Stdout
	logger = log.New(out, "kuler ", log.LstdFlags)
}

func logE(format string, v ...interface{}) {
	if logger == nil {
		return
	}
	logger.Printf("E:"+format+"\n", v...)
}

func logD(format string, v ...interface{}) {
	if logger == nil {
		return
	}
	logger.Printf("D:"+format+"\n", v...)
}
