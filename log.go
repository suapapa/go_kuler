// Copyright 2013, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kuler

import (
	"io"
	"log"
)

var logger *log.Logger

// Set kurler's log out.
// For example, to log on stdout;
//
//     kuler.LogTo(os.Stdout)
//
func LogTo(out io.Writer) {
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
