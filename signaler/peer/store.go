package peer

import (
	"encoding/json"

	"golang.org/x/net/websocket"
)

type PeerStore struct {
	peers []*Peer
}

func (cs *PeerStore) AddFromWs(c *websocket.Conn) error {

	return nil
}

func (cs *PeerStore) Add(c *Peer) {
	cs.peers = append(cs.peers, c)
}

func (cs *PeerStore) Remove(c *Peer) {
	var idx int
	for i, v := range cs.peers {
		if v == c {
			idx = i
		}
	}

	cs.peers = append(cs.peers[:idx], cs.peers[idx+1:]...)
}

func (cs *PeerStore) Notify() {
	for _, v := range cs.peers {
		b, _ := json.Marshal(cs.peers)
		websocket.Message.Send(v.con, string(b))
	}
}
