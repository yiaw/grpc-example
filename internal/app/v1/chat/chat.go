package chat

import (
	"context"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	emptypb "github.com/golang/protobuf/ptypes/empty"
	chatpb "github.com/yiaw/grpc-example/protos/v1/chat"
)

/*

type ChatServiceServer interface {
	SendMessage(context.Context, *Message) (*emptypb.Empty, error)
	RegistRouterChannel(*emptypb.Empty, ChatService_RegistRouterChannelServer) error
	SendAny(context.Context, *AnyMessage) (*emptypb.Empty, error)
	RegistAnyRouterChannel(*emptypb.Empty, ChatService_RegistAnyRouterChannelServer) error
	mustEmbedUnimplementedChatServiceServer()
}


type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

type AnyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User    string     `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	MsgType int32      `protobuf:"varint,2,opt,name=msgType,proto3" json:"msgType,omitempty"`
	Msg     *anypb.Any `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}
*/
type ChatServer struct {
	chatpb.ChatServiceServer
}

func NewChatServer(s *grpc.Server) *grpc.Server {
	if s == nil {
		return nil
	}

	chatpb.RegisterChatServiceServer(s, &ChatServer{})
	return s
}

func (c *ChatServer) RegistRouterChannel(stream chatpb.ChatService_RegistRouterChannelServer) error {
	// Channel init or Register

	msg, err := stream.Recv()
	if err != nil {
		if err == io.EOF {
			log.Printf("Disconnection Stream %s\n", "unkown")
		} else {
			log.Printf("stream.Recv() Fail err %v\n", err)
		}
	}
	user := msg.User

	recv := make(chan *RecvStream)
	go func() {
		//*chatpb.Message, error
		msg, err := stream.Recv()
		recv <- &RecvStream{msg: msg, err: err}
	}()

	channel, err := Client().InitChannel(user)
	if err != nil {
		return status.Errorf(codes.Internal, "init channel fail err=%v", err)

	}
	log.Printf("Enter Chat Server User %s\n", user)
	for {
		timeout := time.After(5 * time.Second)
		select {
		case rcvmsg := <-recv:
			if rcvmsg.err != nil {
				if rcvmsg.err == io.EOF {
					log.Printf("Disconnection Stream %s\n", user)
				}
			}
			log.Printf("Unkown Message\n")
		case msg := <-channel:
			if msg == nil {
				continue
			}

			if err := stream.Send(msg); err != nil {
				log.Printf("Stream Message Send Fail %v\n", err.Error())
			}
		case <-timeout:
			log.Printf("Stream Message TimeOut, %s\n", user)
			log.Printf("Disconnect Check, %s\n", user)
			ping := &chatpb.Message{User: "owner", Msg: "ping message"}
			if err := stream.Send(ping); err != nil {
				log.Printf("Stream Message Send Fail %v\n", err.Error())
				return err
			}
			// NOT TODO
		}
	}
	return nil
}

func (c *ChatServer) SendMessage(ctx context.Context, msg *chatpb.Message) (*emptypb.Empty, error) {
	if err := Client().WriteChannel(msg); err != nil {
		log.Printf("message Send Fail %v\n", err)
		return nil, status.Errorf(codes.Internal, "message send fail %v", err)
	}
	return &emptypb.Empty{}, nil
}

/*
func (c *ChatServer) RegistAnyRouterChannel(user *chatpb.User, stream ChatService_RegistAnyRouterChannelServer) error {
	return nil
}

func (c *ChatServer) SendAny(ctx context.Context, msg *chatpb.AnyMessage) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
*/
