# TCP Chatting with Go

How to use:

Run the Go server:

```
go run .
```

Connect with a client:

```
telnet localhost 8000
```

Connect with another client:

```
telnet localhost 8000
```

Commands:

```
/name <username> - Set a username for the chat
/rooms - List all available rooms
/join <name> - Join one of the rooms or create one
/send <message> - Send a message to everyone in the room
/quit - Quit the TCP chat
```
