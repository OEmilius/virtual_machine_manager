package common

import (
	//	"bytes"
	"log"
	"math/rand" //for test
	"net/http"
	libvirt "virtual_machine_manager/common/libvirt_mock"
)

type Server struct {
	Id     string
	IpPort string
	//Ip                   string
	Url                  string
	Mysql_connect_string string //for save task results
	Mashine              map[string]libvirt.Mashine
}

type Answer struct {
	ComputeServerID string
	MashineID       string //TODO create this ID unique in all system example UUID
	Body            string
}

type Args struct {
	TaskID    int
	MashineID string
	Version   string //example api server will be changed and we will control version
}

func (srv *Server) NewMashine(r *http.Request, args *Args, ans *Answer) error {
	log.Println("New Virt Mashine")
	go libvirt.NewMashine()
	// after finish write to task db that task is finished
	*ans = Answer{ComputeServerID: srv.Id,
		MashineID: RandStringBytes(8), //in future will be gen UUID
		Body:      "New Virt Mashine"}
	log.Println("New_id=", ans.MashineID)
	return nil
}

func (srv *Server) StartMashine(r *http.Request, args *Args, ans *Answer) error {
	log.Println("Starting Virt Mashine", args.MashineID)
	go libvirt.StartMashine()
	// after finish write to task db that task is finished
	*ans = Answer{ComputeServerID: srv.Id,
		MashineID: args.MashineID,
		Body:      "Starting Virt Mashine"}
	return nil
}

func (srv *Server) StopMashine(r *http.Request, args *Args, ans *Answer) error {
	log.Println("Stoping Virt Mashine", args.MashineID)
	go libvirt.StopMashine()
	// after finish write to task db that task is finished
	*ans = Answer{ComputeServerID: srv.Id,
		MashineID: args.MashineID,
		Body:      "Stoping Virt Mashine"}
	return nil
}

func (srv *Server) DelMashine(r *http.Request, args *Args, ans *Answer) error {
	log.Println("Deleting Virt Mashine", args.MashineID)
	go libvirt.DelMashine()
	// after finish write to task db that task is finished
	*ans = Answer{ComputeServerID: srv.Id,
		MashineID: args.MashineID,
		Body:      "Deleting Virt Mashine"}
	return nil
}

//TODO
// method for check state Compute server

// ********  during Reserch
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
