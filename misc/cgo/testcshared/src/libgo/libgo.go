// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	_ "p"
	"syscall"
	"time"
)

import "C"

var initCh = make(chan int, 1)
var ranMain bool

func init() {
	os.Exit(1)
	fmt.Fprintln(os.Stderr, "In init now")
	// emulate an exceedingly slow package initialization function
	time.Sleep(100 * time.Millisecond)
	initCh <- 42
}

func main() {
	os.Exit(2)
	ranMain = true
}

//export DidInitRun
func DidInitRun() bool {
	fmt.Fprintln(os.Stderr, "IN DID INIT RUN")
	select {
	case x := <-initCh:
		if x != 42 {
			// Just in case initCh was not correctly made.
			println("want init value of 42, got: ", x)
			syscall.Exit(2)
		}
		return true
	default:
		return false
	}
}

//export DidMainRun
func DidMainRun() bool {
	return ranMain
}
