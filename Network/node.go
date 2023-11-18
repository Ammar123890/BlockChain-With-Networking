package network

import (
	"encoding/json"
	"log"

	MidLevelBlockchain "github.com/Ammar123890/Mid-Level-Blockchain/MidLevelBlockchain"
)

var knownNodes = []string{"localhost:8001", "localhost:8002"} // example addresses
type Message struct {
	Type string // e.g., "NewBlock", "NewTransaction"
	Data []byte // Encoded data (block, transaction, etc.)
}

/*
 * @createdby: Syed Muhammad Ammar
 * @description: This function is used to send the message to the node
 * @param: address of the node and instance of message
 * @return: nil
 */
func (n *Node) BroadcastNewBlock(block *MidLevelBlockchain.Block) {
	encodedBlock, err := json.Marshal(block) // Encoding the block
	if err != nil {
		log.Println("Error encoding block:", err)
		return
	}

	msg := &Message{
		Type: "NewBlock",
		Data: encodedBlock,
	}

	for _, nodeAddr := range knownNodes {
		if nodeAddr != n.Address { // Avoid sending it to itself
			go n.sendMessage(nodeAddr, msg) // Send the encoded block to each known node
		}
	}
}

/*
 * @createdby: Syed Muhammad Ammar
 * @description: This function is used to validate the block that is received from the node
 * @param: instance of block
 */
func (n *Node) validateBlock(block *MidLevelBlockchain.Block) bool {

	calculatedHash := block.CalculateHash()
	// Get the last block's hash from the current blockchain for comparison
	var lastBlockHash string
	if len(n.Blockchain.Blocks) > 0 {
		lastBlockHash = n.Blockchain.Blocks[len(n.Blockchain.Blocks)-1].CurrentHash
	}
	return block.CurrentHash == calculatedHash && block.PreviousHash == lastBlockHash
}

/*
 * @createdby: Syed Muhammad Ammar
 * @description: This function is used to send the message to the node
 * @param: address of the node and instance of message
 */

func (n *Node) handleMessage(msg *Message) {
	switch msg.Type {
	case "NewBlock":
		var block MidLevelBlockchain.Block
		err := json.Unmarshal(msg.Data, &block)
		if err != nil {
			log.Println("Error decoding block:", err)
			return
		}
		// Validate the block's hash and the previous hash
		log.Println("Validating block...")
		if n.validateBlock(&block) {
			n.Blockchain.Blocks = append(n.Blockchain.Blocks, &block) // Add the block to the blockchain
			log.Println("New block added")
		} else {
			log.Println("Invalid block received or previous hash does not match")
		}

		// For future implemenation into depth like transation cache and block cache
	}
}

/*
 * @createdby: Syed Muhammad Ammar
 * @description: This function is used to send the broadcast the new transaction to the node
 * @param: address of the node and instance of message
 */

func (n *Node) BroadcastNewTransaction(transaction string) {
	encodedTx, err := json.Marshal(transaction) // Encoding the transaction
	if err != nil {
		log.Println("Error encoding transaction:", err)
		return
	}

	msg := &Message{
		Type: "NewTransaction",
		Data: encodedTx,
	}

	for _, nodeAddr := range knownNodes {
		if nodeAddr != n.Address { // Avoid sending it to itself
			go n.sendMessage(nodeAddr, msg)
		}
	}
}
