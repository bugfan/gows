package impl

import (
	"errors"
	"sync"

	"github.com/gorilla/websocket"
)

type Connection struct {
	wsConn    *websocket.Conn
	inChan    chan []byte
	outChan   chan []byte
	closeChan chan byte
	mutex     sync.Mutex
	isClosed  bool
}

const (
	len = 1000
)

func InitConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		wsConn:    wsConn,
		inChan:    make(chan []byte, len),
		outChan:   make(chan []byte, len),
		closeChan: make(chan byte, 1),
	}
	go conn.readLoop()
	go conn.writeLoop()
	return
}

func (c *Connection) ReadMessage() (data []byte, err error) {
	select {
	case data = <-c.inChan:
	case <-c.closeChan:
		err = errors.New("connection r is closed!")
	}
	return
}
func (c *Connection) WriteMessage(data []byte) (err error) {
	select {
	case c.outChan <- data:
	case <-c.closeChan:
		err = errors.New("connection w is closed!")
	}
	return
}
func (c *Connection) Close() {
	c.wsConn.Close()
	//exec once
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.isClosed != true {
		close(c.closeChan)
		c.isClosed = true
	}
}
func (c *Connection) readLoop() {
	var (
		data []byte
		err  error
	)
	for {
		if _, data, err = c.wsConn.ReadMessage(); err != nil {
			c.Close()
			return
		}
		select {
		case c.inChan <- data:
		case <-c.closeChan:
			c.Close()
		}
	}
}
func (c *Connection) writeLoop() {
	var (
		data []byte
		err  error
	)
	for {
		select {
		case data = <-c.outChan:
		case <-c.closeChan:
			c.Close()
		}
		if err = c.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
			c.Close()
			return
		}
	}
}
