package main

import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
	"time"
)

var ConnectionStatus bool
var ServiceStatus bool

func receiveData(s net.Conn) {
	for {
		message, err := bufio.NewReader(s).ReadString('\n')
		if err != nil {
			fmt.Println("Lost connection with server")
			s.Close()
			ConnectionStatus = false
			return
		}
		fmt.Print(message)
	}
}

func sendData(s net.Conn) {
	for {
		_, err := fmt.Fprintf(s, "Client is alive! \n")
		if err != nil {
			return
		}
		time.Sleep(time.Duration(3) * time.Second)
	}
}

func WorkMonitor(arr []string) {
	for i := 0; i < len(arr); i++ {
		cmd := exec.Command(arr[i])
		err := cmd.Run()
		if err != nil {
			fmt.Println("Can't execute command ", err)
			continue
		}
	}
	fmt.Println("sevice(s) activated")
}

func main() {
	fmt.Println("Client started")
	CommandList := []string{"systemctl start nginx.service"}  //This array contains commands that needs to be exicuted on this node
	DeactivateCommandList := []string{"systemctl stop nginx"} //Deactivation commands for commands in CommandList
	for {
		CurrentConnection := ConnectionStatus
		if !CurrentConnection {
			s, err := net.Dial("tcp", "127.0.0.1:8068") //protocol and IP adr for connection
			if err != nil {
				fmt.Println(err.Error())
				WorkMonitor(CommandList)                    // Activate service and host is until connection is restored
				time.Sleep(time.Duration(10) * time.Second) // seconds before reconnecting
				continue
			}
			fmt.Println(s.RemoteAddr().String() + ": connected")
			ConnectionStatus = true
			WorkMonitor(DeactivateCommandList)
			go sendData(s)
			go receiveData(s)
		}
		time.Sleep(time.Duration(3) * time.Second)
	}
}
