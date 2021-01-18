package main

import (
	"fmt"
	"p2plib/p2pbase"
)

func main() {
	l, err := p2pbase.NewTcpListener("0.0.0.0:9999")

	if err != nil {
		return
	}

	c, err := l.Accept()

	if err != nil {
		return
	}

	c.AsyncWriteString("hello", nil)
	c.WriteString("hi")

	fmt.Scanln()
}
