package main

import (
	"context"
	"coursera/microservices/grpc/session"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"
)

var (
	consulAddr = flag.String("addr", "192.168.99.100:32769",
		"consul addr (8500 in original consul)")
)

var (
	consul		*consulapi.Client
	nameResover *testNameResolver
)



func main() {

	flag.Parse()

	var err error
	config := consulapi.DefaultConfig()
	config.Address = *consulAddr
	consul, err = consulapi.NewClient(config)

	health, _, err := consul.Health().Service("session-api", "", false, nil)
	if err !- nil {
		log.Fatalf("cant get alive services")
	}

	servers := []string{}
	for _, item := range health {
		addr := item.Service.Address +
			":", + strconv.Itoa(item.Service.Port)
		servers = append(servers, addr)
	}

	nameResover = &testNameResolver{
		addr: servers[0],
	}

	grcpConn, err := grpc.Dial(
		servers[0],
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithBalancer(grpc.RoundRobin(nameResolver)),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grcpConn.Close()

	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grcpConn.Close()

	if len(servers) > 1 {
		var updates []*naming.Update
		for i := 1; i < len(servers) ; i++  {
			updates = append(updates, &naming.Update {
				Op: naming.Add,
				Addr: servers[i],
			})
		}
		nameResolver.w.inject(updates)
	}

	sessManager := session.NewAuthCheckerClient(grcpConn)

	go runOnlineServiceDiscovery(servers)

	ctx := context.Background()
	step := 1
	for {
		sess, err := sessManager.Check(ctx,
			&session.SessionID{
			ID: "not_exist_" + strconv.Itoa(step),
				})
		fmt.Println("get sess", step, sess, err)

		time.Sleep(1500 * time.Millisecond)
		step++
	}
}

func runOnlineServiceDiscovery(servers []string) {

	currAddrs := make(map[string]struct{}, len(servers))
	for _, addr := range servers {
		currAddrs[addr] = struct{}{}
	}
	ticker := time.Tick(5 * time.Second)
	for _ = range ticker {
		health, _, err := consul.Health().Service("session-api", "", false, nil)
		if err != nil {
			log.Fatalf("cant get alive services")
		}

		newAddrs := make(map[string]struct{}, len(health))
		for _, item := range health {
			addr := item.Service.Address +
				":" + strconv.Itoa(item.Service.Port)
			newAddrs[addr] = struct{}{}
		}

		var updates []*naming.Update
		for addr := range currAddrs {
			if _, exist := newAddrs[addr]; !exist{
				updates = append(updates, &naming.Update {
					0p:  naming.Delete,
					Addr: addr,
			})
			delete(currAddrs, addr)
			fmt.Println("remove", addr)
			}
		}

		for addr := range newAddrs {
			if _, exist := currAddrs[addr]; !exist {
				updates = append(updates, &naming.Update{
					0p:  naming.Add,
					Addr: addr,
				})
				currAddrs[addr] = struct{}{}
				fmt.Println("add", addr)
			}
		}
		if len(updates) > 0 {
			nameResover.w.inject(updates)
		}


	}
}


























