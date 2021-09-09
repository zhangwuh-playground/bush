package main

import (
	"Bush/gen-go/user_service"
	"Bush/log"
	"Bush/tracing"
	"context"
	"encoding/json"
	"fmt"

	"github.com/opentracing/opentracing-go"
)

type UserService struct{}

func (this *UserService) GetUser(ctx context.Context, tc *user_service.TraceContext, id int32) (*user_service.RcpResponse, error) {
	carrier := opentracing.TextMapCarrier{}
	if err := json.Unmarshal([]byte(tc.GetCarrier()), &carrier); err != nil {
		log.ErrorNt("err when get carrier from trace context", err)
	}

	pspan, err := tracing.GetTracer().Extract(opentracing.TextMap, carrier)
	if err != nil {
		log.ErrorNt("err when extract trace context", err)
	}
	span := tracing.GetTracer().StartSpan("Get user in bush", opentracing.SpanReference{
		Type: opentracing.ChildOfRef, ReferencedContext: pspan,
	})

	defer span.Finish()
	return &user_service.RcpResponse{
		UserInfo: &user_service.UserInfo{
			ID: id, Name: fmt.Sprintf("user:%d", id),
		},
	}, nil
}
