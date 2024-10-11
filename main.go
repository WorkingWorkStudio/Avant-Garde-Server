package main

import (
	"fmt"
	"net"
)
func main(){
	//creates the server
	listner,err:=net.Listen("tcp","127.0.0.1:8000")
	if(err!=nil){
		fmt.Println("Error with server : ",err.Error())
		return
	}
	defer listner.Close()
	//While loop that keeps listning for the server (runs forever)
	for{
		conn,err:=listner.Accept()
		if(err!=nil){
			fmt.Println("Error with server",err.Error())
		}
		go handleConnection(conn)
	}
}
func handleConnection(conn net.Conn){
	defer conn.Close()
	//read the data
	buf:=make([]byte, 1024)
	conn.Read(buf)
	conn.Write([]byte("Message received."))
}