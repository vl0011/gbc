package main

import (
	"p2plib/p2pbase"
)

func main() {
	c, err := p2pbase.NewTcpConnect("0.0.0.0:9999")

	if err != nil {
		return
	}

	s, _ := c.ReadString()
	println(s)

}
