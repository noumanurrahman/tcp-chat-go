package chat

type commandId int

const (
	CMD_NAME commandId = iota
	CMD_JOIN
	CMD_ROOMS
	CMD_SEND
	CMD_QUIT
)

type command struct {
	id     commandId
	client *client
	args   []string
}
