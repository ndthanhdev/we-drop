package boot

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ndthanhdev/we-sync/peer"
	"golang.org/x/net/websocket"
)

var store *peer.PeerStore = &peer.PeerStore{}

func wsHandler(ctx echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		var err error

		cl := peer.New("c1", ws)

		store.Add(cl)
		defer ws.Close()
		defer store.Remove(cl)

		for {
			// Read
			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				ctx.Logger().Error(err)
			}
			fmt.Printf("%s\n", msg)
			cl.HandleMessage(msg)
			store.Notify()
		}
	}).ServeHTTP(ctx.Response(), ctx.Request())
	return nil
}

func Boot() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "../public")
	e.GET("/ws", wsHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
