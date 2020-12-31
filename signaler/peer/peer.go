package peer

import (
	"golang.org/x/net/websocket"
)

type Peer struct {
	Name string
	con  *websocket.Conn
	Locs []string
}

func New(name string, con *websocket.Conn) *Peer {
	return &Peer{
		Name: name,
		con:  con,
		Locs: []string{},
	}
}

func (c *Peer) AddLoc(l string) {
	c.Locs = append(c.Locs, l)
}

func (c *Peer) RemoveLoc(l string) {
	var i int
	for idx, v := range c.Locs {
		if l == v {
			i = idx
		}
	}

	c.Locs = append(c.Locs[:i], c.Locs[i+1:]...)
}
