package main

import "fmt"
import "net"
import "strings"

var myPort string
var asPeer = false

var hostInfo string
var hostAddr *net.UDPAddr
var hostConn *net.UDPConn

var peerInfo string
var peerAddr *net.UDPAddr
var peerConn *net.UDPConn

func initHost() {
  hostInfo = fmt.Sprintf(":%s", myPort)
  hostAddr, err := net.ResolveUDPAddr("udp", hostInfo)
  panicIfError(err)

  hostConn, err = net.ListenUDP("udp", hostAddr)
  panicIfError(err)

  go func() {
    for {
      readCmd()
    }
  }()
}

func closeHost() {
  hostConn.Close()
}

func connectPeerIfSet() {
  if peerInfo != "" {
    peerAddr, err := net.ResolveUDPAddr("udp", peerInfo)
    panicIfError(err)
    local, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
    panicIfError(err)
    peerConn, err = net.DialUDP("udp", local, peerAddr)
    panicIfError(err)
  }
}

func closePeer() {
  if peerInfo != "" {
    peerConn.Close()
  }
}

func pingPeer() {
  if peerInfo != "" {
    sendCmd("ping", fmt.Sprintf(":%s", myPort))
  }
}

func pongPeer() {
  if peerInfo != "" {
    sendCmd("pong", "")
  }
}

func readCmd() {
  buf := make([]byte, 1024)
  n, from, err := hostConn.ReadFromUDP(buf)
  if err == nil {
    msg := string(buf[0:n])
    msgs := strings.Split(msg, " ")
    cmd := msgs[0]
    body := msgs[1]

    setLog(fmt.Sprintf("From %s - Cmd: %s, Body: %s", from, cmd, body))
    switch cmd {
    case "ping":
      ip := strings.Split(from.String(), ":")[0]
      port := strings.Split(body, ":")[1]

      peerInfo = fmt.Sprintf("%s:%s", ip, port)
      connectPeerIfSet()
      pongPeer()
    }
  }
}

func sendCmd(cmd, body string) {
  buf := []byte(fmt.Sprintf("%s %s", cmd, body))
  _, _ = peerConn.Write(buf)
}
