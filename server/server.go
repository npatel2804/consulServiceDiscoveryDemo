package main

import (
	"context"
	"fmt"
	"net"

	api "github.com/hashicorp/consul/api"
	pb "github.com/npatel2804/consulServiceDiscoveryDemo/server/protos"
	gRPC "google.golang.org/grpc"
)

type service struct{}

func (s *service) SumNum(ctx context.Context, r *pb.FunctionRequest) (*pb.FunctionResponse, error) {
	result := r.Num1 - r.Num2
	fmt.Println(result)
	//log.Fatalf(string(result))
	return &pb.FunctionResponse{Result: result}, nil
}

func (s *service) SubNum(ctx context.Context, r *pb.FunctionRequest) (*pb.FunctionResponse, error) {
	result := r.Num1 - r.Num2
	return &pb.FunctionResponse{Result: result}, nil
}

func main() {
	lst, err := net.Listen("tcp", ":50051")
	fmt.Println("starting Server....")
	fmt.Println(err)
	grpc := gRPC.NewServer()
	fmt.Println("server is created")
	pb.RegisterCalculatorServer(grpc, &service{})
	//Here running agent will be pointed by api.DefaultConfig()
	clnt, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		fmt.Printf("%v", err)
	}
	agnt := clnt.Agent()
	sevc := api.AgentServiceRegistration{Name: "SubNum"}
	err1 := agnt.ServiceRegister(&sevc) //it will create object of consul
	if err1 != nil {
		fmt.Printf("%v", err1)
	}
	//the run above command we have to run consule agent in othere terminal
	//for more info visit : https://www.consul.io/docs/agent/basics.html

	//if err1 := grpc.Serve(lst)
	//	ifaces, err := net.Interfaces()
	// handle err
	/*for _, i := range ifaces {
		addrs, _ := i.Addrs()
		// handle err

		var ip net.IP
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			// process IP address
		}
		fmt.Println(ip)
	}*/
	if err1 := grpc.Serve(lst); err1 != nil {
		fmt.Printf("There is error in Serving %v", err1)
	}
	//fmt.Println(err1)

}
