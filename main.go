package main

import (
	"Bush/gen-go/user_service"
	"Bush/log"
	"Bush/tracing"

	"github.com/apache/thrift/lib/go/thrift"
)

const addr = ":9090"

func main() {
	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(&thrift.TConfiguration{})
	transportFactory := thrift.NewTTransportFactory()

	//handler
	handler := &UserService{}

	//transport,no secure
	var err error
	var transport thrift.TServerTransport
	transport, err = thrift.NewTServerSocket(addr)
	if err != nil {
		log.ErrorNt("error running server", err)
	}
	closer := tracing.InitJaeger()
	defer closer.Close()
	//processor
	processor := user_service.NewUserServiceProcessor(handler)

	log.InfoNt(log.Message("Buth comming from %s", addr))

	//start tcp server
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	err = server.Serve()

	if err != nil {
		log.ErrorNt("error running server", err)
	}
}
