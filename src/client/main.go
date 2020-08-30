package main

import (
	"fmt"
	"os"
	"net"
	"bufio"
	"sync"
)

var remoteAddress = "192.168.31.38:8888"

var wg sync.WaitGroup

func clientError(err error, when string)  {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}

func main()  {
	conn, err := net.Dial("tcp",remoteAddress)
	clientError(err, "client conn error")

	buffer := make([]byte, 1024)

	fmt.Println("欢迎来到聊天室！下面可以畅所欲言了！")
	wg.Add(2)
	go sendMsg(conn)
	go readMsg(conn, buffer)
	wg.Wait()
}

func readMsg(conn net.Conn, buffer []byte)  {
	for {
		n,_ := conn.Read(buffer)
		//clientError(err, "client read error")
		serverMsg := string(buffer[0:n])
		if serverMsg != "" {
			fmt.Printf("%s\n", serverMsg)
		}
	}
	wg.Done()
}

func sendMsg(conn net.Conn)  {
	for {
		reader := bufio.NewReader(os.Stdin)
		lineBytes, _, _ := reader.ReadLine()
		conn.Write(lineBytes)
	}
	wg.Done()
}