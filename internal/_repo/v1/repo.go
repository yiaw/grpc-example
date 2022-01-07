package repo

import (
	v1userpb "grpc-example/protos/v1/user"
	v2userpb "grpc-example/protos/v2/user"
)

type UserRepo struct {
	UserId             string
	Password           string
	Labelkey           string
	Labelvalue         string
	PasswordExpiretime string
	Maxuser            int32
	Locallogin         bool
}

var user map[string]*UserRepo

func init() {
	user = make(map[string]*UserRepo)
}

func MapperV1User(u *v1userpb.UserProto) *UserRepo {
	return &UserRepo{
		UserId:             u.UserId,
		Password:           u.Password,
		Labelkey:           u.Labelkey,
		Labelvalue:         u.Labelvalue,
		PasswordExpiretime: u.PasswordExpiretime,
		Maxuser:            u.Maxuser,
		Locallogin:         u.Locallogin,
	}
}

func MapperV2User(u *v2userpb.UserProto) *UserRepo {
	return &UserRepo{
		UserId:             u.UserId,
		Password:           u.Password,
		Labelkey:           u.Labelkey,
		Labelvalue:         u.Labelvalue,
		PasswordExpiretime: u.PasswordExpiretime,
		Maxuser:            u.Maxuser,
		Locallogin:         u.Locallogin,
	}
}

func (u *UserRepo) ConvertV1User() *v1userpb.UserProto {
	return &v1userpb.UserProto{
		UserId:             u.UserId,
		Password:           u.Password,
		Labelkey:           u.Labelkey,
		Labelvalue:         u.Labelvalue,
		PasswordExpiretime: u.PasswordExpiretime,
		Maxuser:            u.Maxuser,
		Locallogin:         u.Locallogin,
	}
}

func (u *UserRepo) ConvertV2User() *v2userpb.UserProto {
	return &v2userpb.UserProto{
		UserId:             u.UserId,
		Password:           u.Password,
		Labelkey:           u.Labelkey,
		Labelvalue:         u.Labelvalue,
		PasswordExpiretime: u.PasswordExpiretime,
		Maxuser:            u.Maxuser,
		Locallogin:         u.Locallogin,
	}

