package network

import (
	"encoding/json"
	"log"
	"net"

	MidLevelBlockchain "github.com/Ammar123890/Mid-Level-Blockchain/MidLevelBlockchain"
)

type Node struct {
	Blockchain *MidLevelBlockchain.Blockchain
	Address    string // Node's network address
	// Additional networking properties will be added later
}

/**
 * @createdby: Syed Muhammad Ammar
 * @description: This function is used to encode the message
 * @param: instance of message in byte slice
 * @return: byte slice and error if any
 **/

func EncodeMessage(msg *Message) ([]byte, error) {
	return json.Marshal(msg)
}

/**
 * @createdby: Syed Muhammad Ammar
 * @description: This function is used to decode the message
 * @param: byte slice of message
 * @return: instance of message and error if any
 **/

func DecodeMessage(data []byte) (*Message, error) {
	var msg Message
	err := json.Unmarshal(data, &msg)
	return &msg, err
}

/**
 * @createdby: Syed Muhammad Ammar
 * @description: This function is used to start the server for now it is listening on the port 8001 and 8002
 **/
func (n *Node) StartServer() {
	ln, err := net.Listen("tcp", n.Address)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go n.handleConnection(conn)
	}
}

/**
 * @createdby: Syed Muhammad Ammar
 * @description: This function is used to handle the connection which is established between the nodes
 * @param: instance of connection
 **/

func (n *Node) handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024) // Adjust buffer size as needed

	length, err := conn.Read(buf)
	if err != nil {
		log.Println("Error reading:", err.Error())
		return
	}

	msg, err := DecodeMessage(buf[:length])
	if err != nil {
		log.Println("Error decoding message:", err)
		return
	}

	n.handleMessage(msg)
}

/**
 * @createdby: Syed Muhammad Ammar
 * @description: This function is used to send the message to the other nodes
 * @param: address of the node and instance of message
 **/

func (n *Node) sendMessage(addr string, msg *Message) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	encodedMsg, err := EncodeMessage(msg)
	if err != nil {
		log.Println("Error encoding message:", err)
		return
	}

	_, err = conn.Write(encodedMsg)
	if err != nil {
		log.Println("Error sending message:", err)
		return
	}
}
