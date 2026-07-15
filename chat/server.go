package chat

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

type server struct {
	rooms    map[string]*room
	commands chan command
}

func NewServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

func (s *server) Run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NAME:
			s.Name(cmd.client, cmd.args)
		case CMD_JOIN:
			s.Join(cmd.client, cmd.args)
		case CMD_ROOMS:
			s.Rooms(cmd.client, cmd.args)
		case CMD_SEND:
			s.Send(cmd.client, cmd.args)
		case CMD_QUIT:
			s.Quit(cmd.client, cmd.args)
		}
	}
}

func (server *server) NewClient(conn net.Conn) {
	log.Printf("New client connected: %s", conn.RemoteAddr().String())

	c := &client{
		conn:     conn,
		name:     "anonymous",
		commands: server.commands,
	}

	c.ReadInput()
}

func (server *server) Name(cl *client, args []string) {
	cl.name = args[1]
	cl.send(fmt.Sprintf("Ok, I'll call you %s from now on!", cl.name))
}

func (server *server) Join(cl *client, args []string) {
	roomName := args[1]
	r, ok := server.rooms[roomName]
	if !ok {
		r = &room{
			name:    roomName,
			members: make(map[net.Addr]*client),
		}
		server.rooms[roomName] = r
	}

	r.members[cl.conn.RemoteAddr()] = cl

	server.quitCurrentRoom(cl)

	cl.room = r

	r.broadcast(cl, fmt.Sprintf("%s has joined the room", cl.name))
	cl.send(fmt.Sprintf("Welcome to %s", r.name))
}

func (server *server) Rooms(cl *client, args []string) {
	var rooms []string
	for name := range server.rooms {
		rooms = append(rooms, name)
	}
	cl.send(fmt.Sprintf("Available rooms: %s", strings.Join(rooms, ", ")))
}

func (server *server) Send(cl *client, args []string) {
	if cl.room == nil {
		cl.err(errors.New("You must join the room first."))
		return
	}

	cl.room.broadcast(cl, fmt.Sprintf("%s: %s", cl.name, strings.Join(args[1:], " ")))
}

func (server *server) Quit(cl *client, args []string) {
	log.Printf("Client disconnected: %s", cl.conn.RemoteAddr().String())
	server.quitCurrentRoom(cl)
	cl.send("Sad to see you go")
	cl.conn.Close()
}

func (server *server) quitCurrentRoom(cl *client) {
	if cl.room != nil {
		delete(cl.room.members, cl.conn.RemoteAddr())
		cl.room.broadcast(cl, fmt.Sprintf("%s has left the room", cl.name))
	}
}
