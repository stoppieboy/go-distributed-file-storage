package p2p

import (
	"fmt"
	"net"
	"sync"
)

//TCPPeer represents the remote node over a TCP established connection
type TCPPeer struct {
	// conn is the underlying connection of the peer
	conn 	 net.Conn

	// if we dial and retrieve a connection => outBound == true
	// if we accept and retrieve a connection => outBound == false
	outBound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer{
	return &TCPPeer{
		conn: 	  conn,
		outBound: outbound,
	}
}

type TCPTransportOpts struct {
	ListenAddr		string
	HandshakeFunc	HandshakeFunc
	Decoder			Decoder
}

type TCPTransport struct {
	TCPTransportOpts
	listener      	net.Listener

	mu				sync.RWMutex
	peers			map[net.Addr]Peer
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport{
	return &TCPTransport{
		TCPTransportOpts: opts,
	}
}

func (t * TCPTransport) ListenAndAccept() error{
	var err error

	t.listener, err = net.Listen("tcp", t.ListenAddr)
	if err != nil {
		return err
	}
	
	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}
		fmt.Printf("new incoming connection: %+v\n", conn)
		go t.handleConn(conn)
	}
}

type Temp struct{}
		
func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	if err := t.HandshakeFunc(peer); err != nil {
		fmt.Printf("TCP handshake error: %s\n", err)
		conn.Close()
		return
	}

	// Read loop
	msg := &Temp{}
	for {
		if err:= t.Decoder.Decode(conn, msg); err != nil{
			fmt.Printf("TCP error: %s\n", err)
			continue
		}
	}
	
}