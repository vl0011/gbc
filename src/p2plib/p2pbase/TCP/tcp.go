package TCP

import (
	"bytes"
	"net"
)

type TcpConnListener struct {
	l net.Listener
	// c net.ListenConfig
}

// TODO:: 이름 변경

func NewTcpListener(addr string) (*TcpConnListener, error) {
	lit, err := net.Listen("tcp", addr)

	if err != nil {
		return nil, err
	}

	ret := TcpConnListener{
		l: lit,
	}

	return &ret, nil
}

func NewTcpConnect(addr string) (*TcpConn, error) {
	c, e := net.Dial("tcp", addr)
	if e != nil {
		return nil, e
	}

	ret := TcpConn{
		c: c,
	}

	return &ret, nil
}

func (tcb *TcpConnListener) Accept() (*TcpConn, error) {
	c, err := tcb.l.Accept()
	var r TcpConn
	r = TcpConn{
		c: c,
	}

	return &r, err
}

type TcpConn struct {
	c net.Conn
}

func (tc *TcpConn) Read(b []byte) (int, error) {
	return tc.c.Read(b)
}

func (tc *TcpConn) Write(b []byte) (int, error) {
	return tc.c.Write(b)
}

func (tc *TcpConn) ReadString() (string, error) {
	var b bytes.Buffer
	_, err := b.ReadFrom(tc.c)
	return b.String(), err
}

// WriteString(s string) error
// s -> 
func (tc *TcpConn) WriteString(s string) error {
	var b bytes.Buffer
	b.WriteString(s)
	_, err := b.WriteTo(tc.c)
	return err
}


// AsyncRead (b []byte, cb func(int, error))
// b ->	 Byte slice to read
// cb -> Callback function, nil unavailable
//    n   -> Read byte len
//	  err -> Read error
// Read values ​​asynchronously
func (tc *TcpConn) AsyncRead(b []byte, cb func(int, error)) {
	go func() {
		i, err := tc.c.Read(b)
		cb(i, err)
	}()
}

// AsyncWrite (b []byte, cb func(n int, err error))
//  b  -> Text to write
//	cb -> Callback function, nil available
//    n   -> Write byte len
//	  err -> Write error
// Write values ​​asynchronously
func (tc *TcpConn) AsyncWrite(b []byte, cb func(int, error)) {
	go func() {
		i, err := tc.c.Write(b)
		if cb != nil {
			cb(i, err)
		}
	}()
}

// AsyncReadString (cb func(s string, err error))
//	cb -> Callback function, nil not available
//    s	  -> Read text
//	  err -> Read error
// Read values ​​asynchronously
func (tc *TcpConn) AsyncReadString(cb func(string, error)) {
	go func() {
		var b bytes.Buffer
		_, err := b.ReadFrom(tc.c)
		cb(b.String(), err)
	}()
}

// AsyncWriteString (s string, cb func(error))
// 	s -> Write Text
// 	cb -> Callback function, nil available
//	  err -> Read error
// Write text asynchronously
func (tc *TcpConn) AsyncWriteString(s string, cb func(error)) {
	go func(d string) {
		var b bytes.Buffer
		b.WriteString(d)
		_, err := b.WriteTo(tc.c)
		if cb != nil {
			cb(err)
		}
	}(s)
}

// End ()
// Close connection
func (tc *TcpConn) End() {
	tc.c.Close()
}
