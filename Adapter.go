package ws

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/night-codes/tokay"
)

// Adapter is slice of Connect instances
type Adapter struct {
	command    string
	connection *Connection
	data       *[]byte
	requestID  uint64
	sent       bool
}

// NewAdapter makes new *Adapter instance
func newAdapter(command string, connection *Connection, data *[]byte, requestID uint64) *Adapter {
	return &Adapter{
		command:    command,
		connection: connection,
		data:       data,
		requestID:  requestID,
	}
}

// Data returns client message
func (a *Adapter) Data() []byte {
	return *a.data
}

// JSONData makes json.Unmarshal for client message
func (a *Adapter) JSONData(obj interface{}) error {
	return json.Unmarshal(*a.data, obj)
}

// StringData return string client message without json ""
func (a *Adapter) StringData() string {
	return strings.Trim(string(*a.data), "\"")
}

// Context returns copy of *tokay.Context
func (a *Adapter) Context() *tokay.Context {
	return a.connection.Context()
}

// RequestID returns adapter.requestID
func (a *Adapter) RequestID() uint64 {
	return a.requestID
}

// Command returns request command
func (a *Adapter) Command() string {
	return a.command
}

// Send message to open connection
func (a *Adapter) Send(message interface{}) error {
	if a.sent {
		return fmt.Errorf("Adaper already sent")
	}
	a.sent = true
	return a.connection.Send(a.command, message, a.requestID)
}

// Connection returns adapter connect instance
func (a *Adapter) Connection() *Connection {
	return a.connection
}

// Close all connects
func (a *Adapter) Close() {
	a.connection.Close()
}

// Subscribers of commands ("command1,command2" etc.)
func (a *Adapter) Subscribers(commands string) Connections {
	return a.connection.Subscribers(commands)
}

// User returns connection user
func (a *Adapter) User() *User {
	return a.connection.User()
}
