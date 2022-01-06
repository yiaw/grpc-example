package user

import (
	"context"
	"fmt"
	"log"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	userpb "github.com/yiaw/grpc-example/protos/v2/user"
)

type userServer struct {
	userpb.UserServer
}

var defUser map[string]*userpb.UserProto

func init() {
	defUser = make(map[string]*userpb.UserProto)
}

func NewUserServer(s *grpc.Server) *grpc.Server {
	if s == nil {
		return nil
	}

	userpb.RegisterUserServer(s, &userServer{})
	return s
}

//	SetUser(context.Context, *UserProto) (*ResponseData, error)
func (u *userServer) SetUser(ctx context.Context, req *userpb.UserProto) (*userpb.ResponseData, error) {
	_, ok := defUser[req.UserId]
	if ok {
		return nil, status.Errorf(codes.AlreadyExists, "already user: %s", req.UserId)
	}

	defUser[req.UserId] = req
	return &userpb.ResponseData{
		ResponseMessage: fmt.Sprintf("%s Create Succ..", req.UserId),
	}, nil
}

//GetUser(context.Context, *UserId) (*UserProto, error)
func (u *userServer) GetUser(ctx context.Context, req *userpb.UserId) (*userpb.UserProto, error) {
	resUser, ok := defUser[req.UserId]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "not found user: %s", req.UserId)
	}
	return resUser, nil
}

//ListUsers(context.Context, *None) (*ListUsersResponse, error)
func (u *userServer) ListUsers(ctx context.Context, req *userpb.None) (*userpb.ListUsersResponse, error) {
	var resUserList []*userpb.UserProto
	for _, v := range defUser {
		resUserList = append(resUserList, v)
	}

	return &userpb.ListUsersResponse{
		Users: resUserList,
	}, nil
}

//UpdateUser(context.Context, *UserProto) (*ResponseData, error)
func (u *userServer) UpdateUser(ctx context.Context, req *userpb.UserProto) (*userpb.ResponseData, error) {
	_, ok := defUser[req.UserId]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "not found user: %s", req.UserId)
	}

	defUser[req.UserId] = req
	return &userpb.ResponseData{
		ResponseMessage: fmt.Sprintf("%s Update Succ..", req.UserId),
	}, nil
}

//DeleteUser(context.Context, *UserId) (*ResponseData, error)
func (u *userServer) DeleteUser(ctx context.Context, req *userpb.UserId) (*userpb.ResponseData, error) {
	delete(defUser, req.UserId)
	_, ok := defUser[req.UserId]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "not found user: %s", req.UserId)
	}

	return &userpb.ResponseData{
		ResponseMessage: fmt.Sprintf("%s Delete succ..", req.UserId),
	}, nil
}
