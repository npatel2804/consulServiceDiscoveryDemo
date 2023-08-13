package main

import (
	"fmt"

	api "github.com/hashicorp/consul/api"
)

//Here we will Discover the service Registered in Server Section
//using consul agent
//Here running agent will be pointed by api.Defaultconfig()
func main() {
	clnt, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		fmt.Printf("%v", err)
	}
	agnt := clnt.Agent()
	//the run further command we have to run consule agent in othere terminal
	//for more info visit : https://www.consul.io/docs/agent/basics.html
	servicesList, err2 := agnt.Services()
	fmt.Printf("%v\n", err2)

	fmt.Printf("%v\n", servicesList) //This will Print the service Register with consul by server Program

}
