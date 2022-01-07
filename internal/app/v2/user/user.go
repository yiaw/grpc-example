package user

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1repo "github.com/yiaw/grpc-example/internal/repo/v1"
	userpb "github.com/yiaw/grpc-example/protos/v2/user"
)

type userServer struct {
	userpb.UserServer
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
	_, ok := v1repo.User[req.UserId]
	if ok {
		return nil, status.Errorf(codes.AlreadyExists, "already user: %s", req.UserId)
	}

	log.Printf("req=%v\n", req)
	log.Printf("reqid=%v\n", req.UserId)

	user := v1repo.MapperV2User(req)
	if user == nil {
		return nil, status.Errorf(codes.Internal, "MapperV2User Fail")
	}

	v1repo.User[req.UserId] = user
	log.Printf("user=%v\n", user)
	return &userpb.ResponseData{
		ResponseMessage: fmt.Sprintf("%s Create Succ..", req.UserId),
	}, nil
}

//GetUser(context.Context, *UserId) (*UserProto, error)
func (u *userServer) GetUser(ctx context.Context, req *userpb.UserId) (*userpb.UserProto, error) {
	resUser, ok := v1repo.User[req.UserId]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "not found user: %s", req.UserId)
	}
	log.Printf("resUser=%v\n", resUser)

	user := v1repo.ConvertV2User(resUser)
	if user == nil {
		return nil, status.Errorf(codes.Internal, "ConvertV2User Fail")
	}
	log.Printf("user=%v\n", user)
	return user, nil
}

//ListUsers(context.Context, *None) (*ListUsersResponse, error)
func (u *userServer) ListUsers(ctx context.Context, req *userpb.None) (*userpb.ListUsersResponse, error) {
	var resUserList []*userpb.UserProto
	for _, v := range v1repo.User {
		resUserList = append(resUserList, v1repo.ConvertV2User(v))
	}

	return &userpb.ListUsersResponse{
		Users: resUserList,
	}, nil
}

//UpdateUser(context.Context, *UserProto) (*ResponseData, error)
func (u *userServer) UpdateUser(ctx context.Context, req *userpb.UserProto) (*userpb.ResponseData, error) {
	_, ok := v1repo.User[req.UserId]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "not found user: %s", req.UserId)
	}

	user := v1repo.MapperV2User(req)
	if user == nil {
		return nil, status.Errorf(codes.Internal, "MapperV2User Fail")
	}

	v1repo.User[req.UserId] = user
	return &userpb.ResponseData{
		ResponseMessage: fmt.Sprintf("%s Update Succ..", req.UserId),
	}, nil
}

//DeleteUser(context.Context, *UserId) (*ResponseData, error)
func (u *userServer) DeleteUser(ctx context.Context, req *userpb.UserId) (*userpb.ResponseData, error) {
	delete(v1repo.User, req.UserId)
	_, ok := v1repo.User[req.UserId]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "not found user: %s", req.UserId)
	}

	return &userpb.ResponseData{
		ResponseMessage: fmt.Sprintf("%s Delete succ..", req.UserId),
	}, nil
}
