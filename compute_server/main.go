/*
read config.json
register new json rpc server
wait commands
Utero in common
*/
package main

import (
	"fmt"
	//"log"
	"net/http"

	"virtual_machine_manager/common"

	libvirt "virtual_machine_manager/common/libvirt_mock"

	cfgdecod "encoding/json"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

//TODO
// get url from file example config.json
// or get url from flag params
// get server id from file or flag params
var cfg Config

type Config struct {
	Id     string
	IpPort string
	//Url     string
	Mysql_connect_string string
}

func main() {
	fmt.Println("start compute server")
	cfg = readConfig("config.json")
	srv := new(common.Server)
	srv.Id = cfg.Id
	//srv.Url = "http://localhost:8000/rpc"
	srv.Mashine = make(map[string]libvirt.Mashine)
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")
	s.RegisterService(srv, "")
	r := mux.NewRouter()
	r.Handle("/rpc", s)
	//go http.ListenAndServe(":8000", r)
	fmt.Println("http.ListenAndServe", cfg.IpPort)
	//TODO check error starting for prevent busy port
	if err := http.ListenAndServe(cfg.IpPort, r); err != nil {
		//fmt.Println(err)
		panic(err)
	}
	//fmt.Scanln()
}

//read config from config.json file
func readConfig(fname string) Config {
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println("opening config file error", fname, err)
		panic(err)
	}
	decoder := cfgdecod.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("error reading config file", fname, err)
		panic(err)
	}
	//fmt.Println("config=", config)
	return config
}
