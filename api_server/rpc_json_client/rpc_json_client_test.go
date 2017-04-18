package rpc_json_client

import (
	"fmt"
)

func ExampleStartVirtMashine() {
	//before run this example need to start Compute server on localhost:8000 port
	srv1 := ComputeServer{}
	srv1.Url = "http://localhost:8000/rpc"
	m_id, err := srv1.NewVirtMashine()
	err = srv1.StartVirtMashine(m_id)
	//	srv1.StopVirtMashine(m_id)
	//	srv1.DelVirtMashine(m_id)
	fmt.Println(err)
	//Output: Starting Virt Mashine
	//<nil>

}
