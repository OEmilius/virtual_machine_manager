package rpc_json_client

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"virtual_machine_manager/common"

	"github.com/gorilla/rpc/json"
)

//var ComputeServers = make(map[string]ComputeServer)

type ComputeServer struct {
	State    string
	Id       string
	Url      string
	Ip       string
	Mashines map[string]bool
}

func (srv *ComputeServer) NewVirtMashine() (id string, err error) {
	args := &common.Args{
		MashineID: "~~~new ~~~"} //who will do this? it will be in param or return param?
	ans, _ := srv.sendCommand("Server.NewMashine", args)
	return ans.MashineID, nil
}

func (srv *ComputeServer) StartVirtMashine(m_id string) (err error) {
	args := &common.Args{
		MashineID: m_id}
	ans, _ := srv.sendCommand("Server.StartMashine", args)
	fmt.Println(ans.Body)
	return nil
}

func (srv *ComputeServer) StopVirtMashine(m_id string) (err error) {
	args := &common.Args{
		MashineID: m_id}
	ans, _ := srv.sendCommand("Server.StopMashine", args)
	fmt.Println(ans.Body)
	return nil
}

func (srv *ComputeServer) DelVirtMashine(m_id string) (err error) {
	args := &common.Args{
		MashineID: m_id}
	ans, _ := srv.sendCommand("Server.DelMashine", args)
	fmt.Println(ans.Body)
	return nil
}

func (srv *ComputeServer) sendCommand(cmd string, args *common.Args) (ans common.Answer, err error) {
	url := srv.Url
	message, err := json.EncodeClientRequest(cmd, args)
	if err != nil {
		log.Fatalf("%s", err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
	if err != nil {
		log.Fatalf("%s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error in sending request to %s. %s", url, err)
	}
	//defer resp.Body.Close()

	//var ans common.Answer
	err = json.DecodeClientResponse(resp.Body, &ans)
	if err != nil {
		log.Fatalf("Couldn't decode response. %s", err)
	}
	resp.Body.Close()
	log.Printf("mashineID =%s, answer=%s\n", ans.MashineID, ans.Body)
	return ans, nil
}

//func main() {
//	fmt.Println("start api server")
//	srv1 := ComputeServer{}
//	srv1.Url = "http://localhost:8000/rpc"
//	m_id, err := srv1.NewVirtMashine()
//	fmt.Println("***new virt mashine id =", m_id, "error=", err)
//	srv1.StartVirtMashine(m_id)
//	srv1.StopVirtMashine(m_id)
//	srv1.DelVirtMashine(m_id)
//}
