package app

import (
	"context"
	"fmt"
	userpb "grpc-example/protos/v1/user"

	"google.golang.org/grpc"
)

type userServer struct {
	userpb.UserServer
}

var defUser map[string]*userpb.UserProto

func init() {
	defUser = make(map[string]*userpb.UserProto)
}
func NewUserServer() *grpc.Server {
	s := grpc.NewServer()
	userpb.RegisterUserServer(s, &userServer{})
	return s
}

//	SetUser(context.Context, *UserProto) (*ResponseData, error)
func (u *userServer) SetUser(ctx context.Context, req *userpb.UserProto) (*userpb.ResponseData, error) {
	defUser[req.UserId] = req
	return &userpb.ResponseData{
		ResponseCode:    200,
		ResponseMessage: fmt.Sprintf("%s Create Succ..", req.UserId),
	}, nil
}

//GetUser(context.Context, *UserId) (*UserProto, error)
func (u *userServer) GetUser(ctx context.Context, req *userpb.UserId) (*userpb.UserProto, error) {
	resUser, ok := defUser[req.UserId]
	if !ok {
		return nil, fmt.Errorf("%s not foun user", req.UserId)
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
	defUser[req.UserId] = req
	return &userpb.ResponseData{
		ResponseCode:    200,
		ResponseMessage: fmt.Sprintf("%s Update Succ..", req.UserId),
	}, nil
}

//DeleteUser(context.Context, *UserId) (*ResponseData, error)
func (u *userServer) DeleteUser(ctx context.Context, req *userpb.UserId) (*userpb.ResponseData, error) {
	delete(defUser, req.UserId)

	return &userpb.ResponseData{
		ResponseCode:    200,
		ResponseMessage: fmt.Sprintf("%s delete succ..", req.UserId),
	}, nil
}
