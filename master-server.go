package main

import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
	"time"
)

var ServerHostname string = "127.0.0.1:8068" // Server IP and Port

func SendData(s net.Conn) {
	counter := 0
	for {
		_, err := fmt.Fprintf(s, "Server is alive! \n") // Send
		counter++
		if err != nil {
			fmt.Println("Message sent!")
			return
		}
		time.Sleep(time.Duration(3) * time.Second) // How much time to wait between heartbeat messages
	}
}
func RecvData(s net.Conn) {
	for {
		MsgBuff, err := bufio.NewReader(s).ReadString('\n')
		if err != nil {
			fmt.Println("Connection with client lost")
			s.Close()
			return
		}
		fmt.Print(MsgBuff)
	}
}
func WorkActivation(arr []string) {
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
	CommandList := []string{"systemctl start nginx.service"} // This array contains commands that need to be executed on this node
	WorkActivation(CommandList)
	SockListner, _ := net.Listen("tcp", ServerHostname) // Choosing protocol and address for connection
	fmt.Println("Server is active and listening on ", ServerHostname)
	fmt.Println("Waiting for connection")
	for {
		s, err := SockListner.Accept() // Accepting connection
		if err != nil {
			fmt.Println("Connection error")
			continue
		}
		fmt.Println(s.RemoteAddr().String() + "  connected")
		go SendData(s) // Handling sending operations
		go RecvData(s) //Handling incoming data
	}
}
