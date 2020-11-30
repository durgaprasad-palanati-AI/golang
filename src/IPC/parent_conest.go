package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os/exec"
	"time"
)

func main() {
	cmd := []string{"go", "run", "conest.go"}

	proc := exec.Command(cmd[0], cmd[1], cmd[2])

	stdin, _ := proc.StdinPipe()
	defer stdin.Close()

	stdout, _ := proc.StdoutPipe()
	defer stdout.Close()

	go func() {
		s := bufio.NewScanner(stdout)
		for s.Scan() {
			fmt.Println("Parent says:" + s.Text())
		}
	}()
	// Start the process
	proc.Start()

	var portnum string

	//portnum := "3333"
	fmt.Println("enter the port number")
	fmt.Scanf("%s", &portnum)
	io.WriteString(stdin, portnum+"\n")
	// 	*******************************************
	time.Sleep(time.Second * 2)

	//listening on port entered
	lis, _ := net.Listen("tcp", ":"+portnum)
	// accept connection
	conn, _ := lis.Accept()

	// run loop some message counter c
	var c int
	c = 3
	for i := 0; i < c; i++ {
		// get message, output
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received:", string(message))

	}
	conn.Close()

	proc.Process.Kill()

}
