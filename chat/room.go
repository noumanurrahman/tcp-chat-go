package chat

import "net"

type room struct {
	name    string
	members map[net.Addr]*client
}

func (room *room) broadcast(sender *client, msg string) {
	for addr, m := range room.members {
		if addr != sender.conn.RemoteAddr() {
			m.send(msg)
		}
	}
}
