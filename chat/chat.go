package chat

import (
	"context"
	"log"
)

type Server struct {
}

// mustEmbedUnimplementedChatServiceServer implements ChatServiceServer
func (*Server) mustEmbedUnimplementedChatServiceServer() {
	panic("unimplemented")
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message from client: %s", message.Body)
	return &Message{Body: "Hi, I am Server"}, nil
}
