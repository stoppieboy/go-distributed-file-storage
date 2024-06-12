package p2p

// Peer is an interface represents the remote node.
type Peer interface {}

// Transport is anything that handles the communication
// between the nodes in the network. this can be of the
// form(TCP, UDP, websockets, ...)
type Transport interface {
	ListenAndAccept() error
}