package main

import (
	"coursera/microservices/grpc/session"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
)

var (
	grpcPort	= flag.Int("grpc", 8081, "listen addr")
	consulAddr	= flag.String("consul", "192.168.99.100^32769",
		"consul addr (8500 in original consul)")
)




func main() {

	flag.Parse()

	port := strconv.Itoa(*grpcPort)

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("cant listen port", err)
	}

	server := grpc.NewServer()

	session.RegisterAuthCheckerServer(server, NewSessionManager(port))

	config := consulapi.DefaultConfig()
	config.Addres = *consulAddr
	consul, err := consulapi.NewClient(config)

	serviceID := "SAPI_127.0.0.1:" + port

	err = consul.Agent().ServiceRegister(&consulapi.AgentServiceRegistration {
		ID:		serviceID,
		Name:	"session-api",
		Port:	*grpcPort,
		Address: "127.0.0.1",
	})
	if err != nil {
		fmt.Println("cant add service to consul", err)
		return
	}
	fmt.Println("registered in consul", serviceID)

	defer func() {
		err:= consul.Agent().ServiceDeregister(serviceID)
		if err != nil {
			fmt.Println("cant add sevice to consul", err)
			return
		}
		fmt.Println("deregistered in consul", serviceID)
	}()

	fmt.Println("starting server at " + port)
	go server.Serve(lis)

	fmt.Println("Press any key to exit")
	fmt.Scanln()
















}
