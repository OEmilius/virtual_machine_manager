package main

import (
	"fmt"
	"virtual_machine_manager/api_server/rest_server"
	"virtual_machine_manager/api_server/rpc_json_client"
	"virtual_machine_manager/cru_task"
	//	"errors"
)

var RpcServers = make(map[int]rpc_json_client.ComputeServer)
var done = make(chan string)

//TODO load config from config.json or find way to add server urls
func main() {
	fmt.Println("start api server")
	co_srv1 := rpc_json_client.ComputeServer{}
	co_srv1.Url = "http://localhost:8000/rpc"
	RpcServers[1] = co_srv1
	co_srv2 := rpc_json_client.ComputeServer{}
	co_srv2.Url = "http://localhost:8001/rpc"
	RpcServers[2] = co_srv2

	cru_task.ChanTask = make(chan cru_task.Task, 20)
	go TaskProcessor(cru_task.ChanTask)
	go rest_server.StartRestServer()
	fmt.Scanln()
	close(cru_task.ChanTask)
	fmt.Println(<-done)
}

func TaskProcessor(in_ch chan cru_task.Task) {
	fmt.Println("Start TaskProcessor")
	fmt.Println("cap=", cap(in_ch))
	//test for one server
	co_srv1 := RpcServers[1]
	for t := range in_ch {
		fmt.Println("****task=", t)
		switch t.Name {
		case "NewVirtMashine":
			fmt.Println(co_srv1.NewVirtMashine())
		case "StartVirtMashine":
			fmt.Println(co_srv1.StartVirtMashine(t.VirtualMashineID))
		case "StopVirtMashine":
			fmt.Println(co_srv1.StopVirtMashine(t.VirtualMashineID))
		case "DelVirtMashine":
			fmt.Println(co_srv1.DelVirtMashine(t.VirtualMashineID))
		}
	}
	done <- "stop task processor"
}

func ChooseServer(m map[int][]string) (int, bool) {
	if len(m) == 0 { // map not have any keys
		return 0, false
	}
	max := int(^uint(0) >> 1)
	var serverNumber int
	for k, v := range m {
		if len(v) < max {
			max = len(v)
			serverNumber = k
		}
	}
	return serverNumber, true
}
