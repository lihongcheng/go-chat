package main

import (
	"net"
	"fmt"
	"os"
	"crypto/md5"
)

var serverAddress = "192.168.31.38:8888"

var msgChannel = make(chan string)

var connPool = make(map[net.Conn]string)

var connPoolChan = make(chan map[net.Conn]string)

var msgConn = make(chan net.Conn)

func main()  {
	listener,err := net.Listen("tcp",serverAddress)
	serverError(err, "net.listen error")

	go broadcast()

	for {
		conn, e := listener.Accept()
		serverError(e, "listener.accept error")

		connPool[conn] = conn.RemoteAddr().String()
		go sayWith(conn)
	}
}

func broadcast()  {

	for {
		msg := <-msgChannel
		if msg != "" {
			connPools := <- connPoolChan
			curConn := <- msgConn
			for conn := range connPools {
				if curConn != conn{
					data := []byte(curConn.RemoteAddr().String())
					has := md5.Sum(data)
					md5str1 := fmt.Sprintf("%x", has)
					conn.Write([]byte("游客"+ md5str1[len(md5str1)-4:] + "说:"+msg))
				}
			}
		}
	}
}


func sayWith(conn net.Conn)  {

	buffer := make([]byte, 1024)
	for {
		n,_ := conn.Read(buffer)
		//serverError(err, "conn.read buffer error")

		clientMsg := string(buffer[0:n])


		data := []byte(conn.RemoteAddr().String())
		has := md5.Sum(data)
		md5str1 := fmt.Sprintf("%x", has)
		user := md5str1[len(md5str1)-4:]

		if clientMsg == "" {
			fmt.Printf("游客%s退出聊天室\n",user)
            delete(connPool, conn)
			break
		}
		fmt.Printf("收到消息，游客%s说：%s\n",user, clientMsg)

		msgChannel <- clientMsg
		connPoolChan <- connPool
		msgConn <- conn

	}
	conn.Close()

}

func serverError(err error, when string)  {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}