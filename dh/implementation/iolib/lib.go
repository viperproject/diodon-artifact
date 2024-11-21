package iolib

import net "net"
import "errors"


const MAX_DATA_SIZE = 1024

type IoLibState struct {
	conn net.Conn
	connClosed bool
}

func NewLibState(endpoint string) (l *IoLibState, err error) {
	conn, err := net.Dial("udp", endpoint)
	if err != nil {
		return nil, err
	}
	state := &IoLibState {
		conn: conn,
		connClosed: false,
	}
	return state, nil
}

func (l *IoLibState) Send(data []byte) (err error) {
	if l.connClosed {
		return errors.New("connection is closed")
	}
	bytesWritten, err := l.conn.Write(data)
	if err != nil {
		return err
	}
	if bytesWritten != len(data) {
		return errors.New("not all data has been sent")
	}
	return nil
}

func (l *IoLibState) Recv() (data []byte, err error) {
	if l.connClosed {
		return nil, errors.New("connection is closed")
	}
	buf := make([]byte, MAX_DATA_SIZE)
	bytesRead, err := l.conn.Read(buf)
	if err != nil {
		return nil, err
	}
	buf = buf[:bytesRead]
	return buf, nil
}

func (l *IoLibState) Close() {
	l.connClosed = true
	l.conn.Close()
}
