package peer

import (
	"encoding/json"
	"fmt"
)

type MessageType string

const (
	AddLoc    MessageType = "add-loc"
	RemoveLoc MessageType = "remove-loc"
)

type Message struct {
	Type MessageType
}

type AddLocMessage struct {
	Loc string
}

type RemoveLocMessage struct {
	Loc string
}

type Parsed struct {
	AddLocMessage    *AddLocMessage
	RemoveLocMessage *RemoveLocMessage
}

func ifAddLoc(t string, msg []byte, addLocMsg *AddLocMessage) (error, *AddLocMessage) {
	if t != string(AddLoc) {
		return nil, nil
	}

	addLocMsg = &AddLocMessage{}
	err := json.Unmarshal(msg, addLocMsg)

	if err != nil {
		addLocMsg = nil
	}

	return err, nil
}

func ifRemoveLoc(t string, msg []byte) (error, *RemoveLocMessage) {
	if t != string(RemoveLoc) {
		return nil, nil
	}

	removeLocMsg := &RemoveLocMessage{}
	err := json.Unmarshal(msg, removeLocMsg)

	if err != nil {
		removeLocMsg = nil
	}

	return err, removeLocMsg
}

func parseMessage(msg string) (*Parsed, error) {
	var message Message
	var bMsg = []byte(msg)
	err := json.Unmarshal(bMsg, &message)

	fmt.Printf("msg %s", msg)

	if err != nil {
		fmt.Printf("error %#v", err)
		return nil, err
	}

	var (
		add    *AddLocMessage
		remove *RemoveLocMessage
	)

	switch message.Type {
	case AddLoc:
		err = json.Unmarshal(bMsg, &add)

	case RemoveLoc:
		err = json.Unmarshal(bMsg, &remove)
	}

	return &Parsed{
		AddLocMessage:    add,
		RemoveLocMessage: remove,
	}, err
}

type Handler struct {
	Type    MessageType
	Handler func(store *PeerStore) (b []byte)
}

type AddLocMsg struct {
	Loc string
}

func (s *AddLocMsg) AddLocHandler(b []byte) error {
	var message AddLocMessage
	err := json.Unmarshal(b, &message)

}

const handlers = []Handler{
	{
		"add-loc",
		AddLocHandler,
	},
}

func (c *Peer) HandleMessage(msg string) error {
	var parsed *Parsed
	parsed, err := parseMessage(msg)

	if err != nil {
		fmt.Printf("%#v", err)
		return err
	}

	fmt.Printf("parsed done")

	if parsed.AddLocMessage != nil {
		c.Locs = append(c.Locs, parsed.AddLocMessage.Loc)
	} else if parsed.RemoveLocMessage != nil {

	}

	return nil
}
