package main 

import ("fmt"
		"log"
		"net"
		"os"
		"bufio"
		"io/ioutil"
		)
func main() {

	dataFromfile,errFile:=ioutil.ReadFile("GoFile.txt")
	if errFile!=nil {
		panic(errFile)
	}

	conn,err:=net.Dial("tcp","127.0.0.1:5000")
	if err!=nil {
		log.Fatalln(err)
	}
	fmt.Println("Connection Established with server ")
	fmt.Printf("%s",dataFromfile) 
	
	
	go WriteToConnection(conn)
	go ReadFromConnection(conn)	
	defer conn.Close()
	for {

	}
	
}

func  WriteToConnection(conn net.Conn) {
	bufioReader := bufio.NewReader(os.Stdin)
	for {
		text,err:= bufioReader.ReadBytes('\n')
		conn.Write(text)
		if string(text)=="exit" || err!=nil {
			if err!=nil {
				log.Fatalln(err)
			}
			return
		}
		
	}
	
}

func ReadFromConnection(conn net.Conn) {
	bufioReader := bufio.NewReader(conn)
	for {
		bytes,err:=bufioReader.ReadBytes('\n')
		if err !=nil {
			log.Fatalln(err)
			return
		}
		fmt.Printf("%s",string(bytes))
	}

}