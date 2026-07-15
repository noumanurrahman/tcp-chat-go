package chat

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type client struct {
	conn     net.Conn
	name     string
	room     *room
	commands chan<- command
}

func (client *client) ReadInput() {
	for {
		msg, err := bufio.NewReader(client.conn).ReadString('\n')
		if err != nil {
			return
		}
		msg = strings.Trim(msg, "\r\n")
		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])
		switch cmd {
		case "/name":
			client.commands <- command{
				id:     CMD_NAME,
				client: client,
				args:   args,
			}
		case "/join":
			client.commands <- command{
				id:     CMD_JOIN,
				client: client,
				args:   args,
			}
		case "/rooms":
			client.commands <- command{
				id:     CMD_ROOMS,
				client: client,
				args:   args,
			}
		case "/send":
			client.commands <- command{
				id:     CMD_SEND,
				client: client,
				args:   args,
			}
		case "/quit":
			client.commands <- command{
				id:     CMD_QUIT,
				client: client,
				args:   args,
			}
		default:
			client.err(fmt.Errorf("unknown command"))
		}
	}
}

func (client *client) err(err error) {
	writeToConn(client.conn, fmt.Sprintf("Error: %s\n", err.Error()))
}

func writeToConn(conn net.Conn, msg string) {
	conn.Write([]byte(msg))
}

func (client *client) send(msg string) {
	writeToConn(client.conn, fmt.Sprintf("> %s\n", msg))
}
