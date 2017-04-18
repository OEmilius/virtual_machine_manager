package cru_task

import (
	//	"database/sql"
	//	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
)

//const PING_DURATION = 90 // интервал пингов MYSQL

//var DB *sql.DB

type Task struct {
	ID               int
	Name             string //GetState, NewMashine, Start, Stop, Del
	ComputeServerID  string
	VirtualMashineID string
	State            string //Accepted, Sended, Running, Finished
	Err              string
}

//var ChanTask = make(chan Task, 20)

var ChanTask chan Task //this chan will create in main according config

var TmpID int

func GetState(id int) string {
	//select State from Task where id = t.ID;
	log.Printf("select State from Task where id = %d;", id)
	return fmt.Sprintf("will be return state from DB for task id=%d", id)
}

func NewTask(t Task) (id int, err error) {
	//insert to db return id
	//insert into Tasks set ComputerServerID=1, VirtualMashineID='asdf', State="new task";
	TmpID++
	t.ID = TmpID
	ChanTask <- t
	log.Println("accpted task", t.Name)
	return TmpID, nil
}

func ChangeState(t Task) error {
	//update into Tasks set ....
	log.Println("update into Tasks set ....")
	return nil
}
