package main

import (
	"fmt"
	"bufio"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		com, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}

}
type client chan<-string
var {
	entering=make(chan client)
	leaving=make(chan client)
	messages=make(chan string)
}
func broadcaster(){
	client:=make(map[client]bool)
	for{
		select{
			case msg:=<-messages:
			for cli:=range client{
				cli<-msg
			}
			case cli:=<-entering:
			clients[cli]=true
			case cli:=<-leaving:
			default(clients,cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn){
	c:=make(chan string)
	go clientWriter(conn,ch)
	who:=conn.RemoteAddr().String()
	ch<-"You are"+who
	messages<-who+"has arrived"
	entering<-ch 
	input:=bufio.NewScanner(conn)
	for input.Scan(){
		messages<-who+":"+input.Text()
	}
	leaving<-ch 
	messages<-who +"has left"
	conn.Close()
}
func clientWriter(conn net.Conn,ch <-chan string){
	for msg:=range ch{
		fmt.Fprintln(conn,msg)
	}
}