package rest_server

import (
	//	"encoding/json"
	"fmt"
	"net/http"

	"log"
	"strconv"
	task "virtual_machine_manager/cru_task"

	"github.com/gorilla/mux"
)

//type Task struct {
//	ID               int    `json:"id"`
//	Name             string //NewMashine, Start, Stop, Del
//	ComputeServerID  string
//	VirtualMashineID string
//	State            string `json:"state"`
//}

var tasks []task.Task

func GetTaskState(w http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL)
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("status code 400 id conversion to int error"))
		return
	}
	//task.GetState(id)
	w.Write([]byte(task.GetState(id)))
	//json.NewEncoder(w).Encode(tasks[id])
	return

}

func NewVirtMashine(w http.ResponseWriter, req *http.Request) {
	//params := mux.Vars(req)
	log.Println(req.Method, req.URL)
	t := task.Task{ //ID: len(tasks),
		Name: "NewVirtMashine",
		//ComputeServerID:  "1",
		//VirtualMashineID: "00",
		State: "Accepted"}
	//tasks = append(tasks, t)
	if task_id, err := task.NewTask(t); err == nil {
		w.Write([]byte(fmt.Sprintf("task_id=%d", task_id)))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("status code 400"))
		w.Write([]byte(fmt.Sprintf("err=%s", err)))
	}
	//json.NewEncoder(w).Encode(t)
	return
}

func StartVirtMashine(w http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL)
	params := mux.Vars(req)
	if m_id, ok := params["m_id"]; ok {
		t := task.Task{ //ID: len(tasks),
			Name: "StartVirtMashine",
			//ComputeServerID:  "1",
			VirtualMashineID: m_id,
			State:            "Accepted"}
		//tasks = append(tasks, t)
		//task.ChanTask <- t
		//json.NewEncoder(w).Encode(t)
		if task_id, err := task.NewTask(t); err == nil {
			w.Write([]byte(fmt.Sprintf("task_id=%d", task_id)))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("status code 400"))
			w.Write([]byte(fmt.Sprintf("err=%s", err)))
		}
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("status code 400 id not found"))
		return
	}
	return
}

func StopVirtMashine(w http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL)
	params := mux.Vars(req)
	if m_id, ok := params["m_id"]; ok {
		t := task.Task{ //ID: len(tasks),
			Name: "StopVirtMashine",
			//ComputeServerID:  "1",
			VirtualMashineID: m_id,
			State:            "Accepted"}
		//		tasks = append(tasks, t)
		//		task.ChanTask <- t
		//		json.NewEncoder(w).Encode(t)
		if task_id, err := task.NewTask(t); err == nil {
			w.Write([]byte(fmt.Sprintf("task_id=%d", task_id)))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("status code 400"))
			w.Write([]byte(fmt.Sprintf("err=%s", err)))
		}
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("status code 400 id not found"))
		return
	}
	return
}

func DelVirtMashine(w http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL)
	params := mux.Vars(req)
	if m_id, ok := params["m_id"]; ok {
		t := task.Task{ //ID: len(tasks),
			Name: "DelVirtMashine",
			//ComputeServerID:  "1",
			VirtualMashineID: m_id,
			State:            "Accepted"}
		//		tasks = append(tasks, t)
		//		task.ChanTask <- t
		//		json.NewEncoder(w).Encode(t)
		if task_id, err := task.NewTask(t); err == nil {
			w.Write([]byte(fmt.Sprintf("task_id=%d", task_id)))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("status code 400"))
			w.Write([]byte(fmt.Sprintf("err=%s", err)))
		}
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("status code 400 id not found"))
		return
	}
	return
}

func StartRestServer() {
	log.Println("start web")
	//task.ChanTask = make(chan task.Task, 20)
	router := mux.NewRouter()
	tasks = append(tasks, task.Task{ID: 0, State: "Accepted"})
	tasks = append(tasks, task.Task{ID: 1, State: "Running"})
	//router.HandleFunc("/task", GetTaskState).Methods("GET")
	router.HandleFunc("/task/getstate/{id}", GetTaskState).Methods("GET")
	//http://127.0.0.1:65534/task/getstate/1
	//TODO change folow methods to POST
	router.HandleFunc("/task/newtask/NewVirtMashine", NewVirtMashine).Methods("GET")
	router.HandleFunc("/task/newtask/StartVirtMashine/{m_id}", StartVirtMashine).Methods("GET")
	router.HandleFunc("/task/newtask/StopVirtMashine/{m_id}", StopVirtMashine).Methods("GET")
	router.HandleFunc("/task/newtask/DelVirtMashine/{m_id}", DelVirtMashine).Methods("GET")
	log.Println("http.ListenAndServe(:65534)")
	if err := http.ListenAndServe(":65534", router); err != nil {
		log.Fatalln(err)
	}
	//fmt.Scanln()
	//	fmt.Println(<-task.ChanTask)
	//	fmt.Println(<-task.ChanTask)
	//	fmt.Println(<-task.ChanTask)
}
