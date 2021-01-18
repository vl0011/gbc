package main

import (
	"eio"
	"io"
	"os"
	"time"
)

func main() {
	f, e := os.Create("/home/dev/proj/gbc/test")
	if e != nil {
		println(e.Error())
		os.Exit(0)
	}
	var w io.Writer
	w = f
	ew := eio.NewEWriter(w)

	ew.AsyncStringWrite("test1\n", nil)
	ew.AsyncStringWrite("test2\n", nil)
	ew.AsyncStringWrite("test3\n", nil)
	ew.AsyncStringWrite("test4\n", nil)
	ew.AsyncStringWrite("test5\n", nil)
	time.Sleep(time.Second * 2)

}
