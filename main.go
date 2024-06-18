package main

import (
	"log"

	"github.com/stoppieboy/go-distributed-file-storage/p2p"
)

func main(){

	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:		":3000",
		Decoder:		p2p.GOBDecoder{},
		HandshakeFunc:	p2p.NOPHandshakeFunc,
	}

	tr := p2p.NewTCPTransport(tcpOpts)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select{}
}