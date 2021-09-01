package main

import (
	"Bush/gen-go/user_service"

	"github.com/apache/thrift/lib/go/thrift"
	log "github.com/cihub/seelog"
)

const addr = ":9090"

func main() {
	initLogger()
	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(&thrift.TConfiguration{})
	transportFactory := thrift.NewTTransportFactory()

	//handler
	handler := &UserService{}

	//transport,no secure
	var err error
	var transport thrift.TServerTransport
	transport, err = thrift.NewTServerSocket(addr)
	if err != nil {
		log.Errorf("error running server:%s", err.Error())
	}

	//processor
	processor := user_service.NewUserServiceProcessor(handler)

	log.Infof("Buth comming from %s", addr)

	//start tcp server
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	err = server.Serve()

	if err != nil {
		log.Infof("error running server:", err)
	}
}

func initLogger() {
	logger, err := log.LoggerFromParamConfigAsFile("logging.xml", nil)
	if err != nil {
		panic(err)
	}
	log.ReplaceLogger(logger)
}