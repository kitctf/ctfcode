//
// Simple network proxy. Allows live inspection and manipulation of packets.
//
// Copyright (c) 2015 Samuel Groß
//

package main

import (
    "os"
    "log"
    "fmt"
    "net"
    "encoding/hex"
)

// Write your own packet handler with this signature
type PacketHandler func (packet []byte) []byte

// Simple packet handler to dump the packets to the console
func dumpPacketHandler(tag string) PacketHandler {
    return func (packet []byte) []byte {
        fmt.Println(tag)
        fmt.Println("------------------------------------------------------")
        fmt.Println(hex.Dump(packet))
        return packet
    }
}


func usage() {
    fmt.Printf("%s <local address> <remote address>\n", os.Args[0])
    fmt.Println("Address format: host:port")
}

// Relay packets from the source to the destination. Call handler for each packet.
func relay(source, dest net.Conn, handler PacketHandler) {
    in := make([]byte, 4096)

    for {
        n, err := source.Read(in)
        if n == 0 || err != nil {
            // source closed on use, pass that on so dest closes as well
            log.Printf("Connection closed: %s\n", source.RemoteAddr().String())
            dest.Close()
            return
        }

        out := handler(in[:n])

        // we can ignore the case where Write() fails since the other goroutines Read() will also fail in that case
        dest.Write(out)
    }
}

func proxy(client, server net.Conn) {
    go relay(client, server, dumpPacketHandler("Client: " + client.RemoteAddr().String()))
    go relay(server, client, dumpPacketHandler("Server: " + server.RemoteAddr().String()))
}

func main() {
    if len(os.Args) < 3 {
        usage()
        os.Exit(0)
    }

    ln, err := net.Listen("tcp", os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Ready. Connect to %s\n", os.Args[1])

    for {
        client, err := ln.Accept()
        if err != nil {
            log.Print(err)
            continue
        }

        server, err := net.Dial("tcp", os.Args[2])
        if err != nil {
            log.Fatal(err)
        }

        log.Printf("New connection from %s\n", client.RemoteAddr().String())
        go proxy(client, server)
    }
}
