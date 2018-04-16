package main 

import ("fmt"
		"net"
		"log"
		"bufio"
		"strconv"
		"os"
		)
type Connection struct {

	connection net.Conn
	index int

	
}
var connections [10]net.Conn
var totalConnections int=0
func main() {
	
	defer CloseConnections()
	f,errFile:=os.OpenFile("GoFile.txt",os.O_APPEND|os.O_WRONLY,0600)
	if errFile!=nil {
		panic(errFile)
	}
	defer f.Close()
	fmt.Printf("Started Server going for the Connections")
	ln,err:=net.Listen("tcp","127.0.0.1:5000")
	if err!=nil {
		log.Fatalln(err)
	}
	for {
		connections[totalConnections],err=ln.Accept()
		if err!=nil {
			log.Fatalln(err)
			break
		}
		fmt.Printf("Connection Established with %d Client ",totalConnections)
		go ProcessConnection(connections[totalConnections],totalConnections,f)
		totalConnections++
	}

}
func CloseConnections() {

	for i:=0; i<totalConnections;i++ {
		connections[i].Close()
	}
	
}
func ProcessConnection(conn net.Conn,index int,fileToWrite *os.File) {
	
	bufioReader:=bufio.NewReader(conn)
		for {
		bytes,err:=bufioReader.ReadBytes('\n')
		if err!=nil {
			log.Fatalln(err)
			return
		}
		fmt.Printf("\n%s",bytes)
		var dataToFile string=strconv.Itoa(index)+":"+string(bytes)
		_,err=fileToWrite.WriteString(dataToFile)
		if err!=nil {
			panic(err)
		}
		fmt.Printf("\nSending data to others")
		Broadcast(bytes,index)



	}
	
}

func Broadcast(bytes []byte,index int) {
	for i:=0 ; i<totalConnections ; i++ {
		if i!=index && connections[i]!=nil {
			var data string = (strconv.Itoa(index))+": "+(string(bytes))
			connections[i].Write([]byte(data));
		}
	}


}	


