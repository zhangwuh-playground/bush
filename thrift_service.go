package main

import (
	"Bush/gen-go/user_service"
	"Bush/log"
	"context"
	"fmt"

)

type UserService struct{}

func (this *UserService) GetUser(ctx context.Context, id int32) (*user_service.RcpResponse, error) {
	log.InfoNt(log.Message("get user by id:%d", id))
	return &user_service.RcpResponse{
		UserInfo: &user_service.UserInfo{
			ID: id, Name: fmt.Sprintf("user:%d", id),
		},
	} , nil
}

