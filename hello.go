// Package websocket implements the WebSocket protocol defined in RFC 6455.
//
// Overview
//
// The Conn type represents a WebSocket connection. A server application calls
// the Upgrader.Upgrade method from an HTTP request handler to get a *Conn:
//
//  var upgrader = websocket.Upgrader{
//      ReadBufferSize:  1024,
//      WriteBufferSize: 1024,
//  }
//
//  func handler(w http.ResponseWriter, r *http.Request) {
//      conn, err := upgrader.Upgrade(w, r, nil)
//      if err != nil {
//          log.Println(err)
//          return
//      }
//      ... Use conn to send and receive messages.
//  }
//
// Call the connection's WriteMessage and ReadMessage methods to send and
// receive messages as a slice of bytes. This snippet of code shows how to echo
// messages using these methods:
//
//  for {
//      messageType, p, err := conn.ReadMessage()
//      if err != nil {
//          log.Println(err)
//          return
//      }
//      if err := conn.WriteMessage(messageType, p); err != nil {
//          log.Println(err)
//          return
//      }
//  }
//
// In above snippet of code, p is a []byte and messageType is an int with value
// websocket.BinaryMessage or websocket.TextMessage.
//
// An application can also send and receive messages using the io.WriteCloser
// and io.Reader interfaces. To send a message, call the connection NextWriter
// method to get an io.WriteCloser, write the message to the writer and close
// the writer when done. To receive a message, call the connection NextReader
// method to get an io.Reader and read until io.EOF is returned. This snippet
// shows how to echo messages using the NextWriter and NextReader methods:
//
//  for {
//      messageType, r, err := conn.NextReader()
//      if err != nil {
//          return
//      }
//      w, err := conn.NextWriter(messageType)
//      if err != nil {
//          return err
//      }
//      if _, err := io.Copy(w, r); err != nil {
//          return err
//      }
//      if err := w.Close(); err != nil {
//          return err
//      }
//  }
//
// Data Messages
//
// The WebSocket protocol distinguishes between text and binary data messages.
// Text messages are interpreted as UTF-8 encoded text. The interpretation of
// binary messages is left to the application.
package awsom_pcakage

import (
	"fmt"
)

//Prints "cool"
func Cool() {
	fmt.Println("cool")
}

//Prints "cool2"
func Cool2() {
	fmt.Println("cool2")

}
