/* @createdby: Syed Muhammad Ammar
 * @StudentId: 20i2417
 * @Assignment: 01
 */

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	MidLevelBlockchain "github.com/Ammar123890/Mid-Level-Blockchain/MidLevelBlockchain"
	network "github.com/Ammar123890/Mid-Level-Blockchain/Network"
)

/**
 * @createdby: Syed Muhammad Ammar
 * @description: This is the main function of the program which is used to run the blockchain application.
 * @param: nil
 **/

func main() {
	port := flag.String("port", "8001", "Port on which the node will listen")
	flag.Parse()

	nodeAddress := "localhost:" + *port

	node := network.Node{
		Blockchain: &MidLevelBlockchain.Blockchain{},
		Address:    nodeAddress,
	}
	go node.StartServer()
	//	blockchain := MidLevelBlockchain.Blockchain{}
	reader := bufio.NewReader(os.Stdin)
	testTransaction := "Sample Transaction Data"
	node.BroadcastNewTransaction(testTransaction)

	for {
		fmt.Println("\nBlockchain Menu:")
		fmt.Println("1. Mine and Broadcast a Block")
		fmt.Println("2. Display Blocks")
		fmt.Println("3. Change a Block's Transaction")
		fmt.Println("4. Verify Blockchain")
		fmt.Println("5. Set number of transactions per block")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid choice.")
			continue
		}

		switch choice {
		case 1:
			mineAndBroadcastBlock(node.Blockchain, &node)
		case 2:
			displayBlocks(node.Blockchain)
		case 3:
			changeBlock(node.Blockchain, reader)
		case 4:
			verifyBlockchain(node.Blockchain)
		case 5:
			setNumberOfTransactionsPerBlock(node.Blockchain, reader)
		case 6:
			fmt.Println("Exiting the blockchain application.")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}

/**
 * @createdby: Syed Muhammad Ammar
 * @description: This function is used to mine and broadcast the block.
 * @param: bc *MidLevelBlockchain.Blockchain, node *network.Node
 * @assumption: The transactions are comma separated.
 **/

func mineAndBroadcastBlock(bc *MidLevelBlockchain.Blockchain, node *network.Node) {
	// Prompt for transactions or gather them from a transaction pool
	fmt.Println("Enter multiple transactions (comma-separated): ")
	transactionsStr, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	transactionsStr = strings.TrimSpace(transactionsStr)

	if len(transactionsStr) == 0 {
		fmt.Println("Transactions cannot be empty.")
		return
	}

	transactions := strings.Split(transactionsStr, ",")
	for i, transaction := range transactions {
		transactions[i] = strings.TrimSpace(transaction)
	}

	previousHash := ""
	if len(bc.Blocks) > 0 {
		previousHash = bc.Blocks[len(bc.Blocks)-1].CurrentHash
	}

	newBlock := bc.MineBlock(transactions, previousHash)
	if newBlock != nil {
		fmt.Println("New block mined successfully. Broadcasting...")
		node.BroadcastNewBlock(newBlock)
	} else {
		fmt.Println("Failed to mine a new block.")
	}

}

/**
 * @createdby: Syed Muhammad Ammar
 * @description: This function is used to display the blocks in a tabular format
 * @param: instance of blockchain
 **/

func displayBlocks(bc *MidLevelBlockchain.Blockchain) {
	fmt.Println("\nBlocks in the blockchain:")
	bc.DisplayBlocks()
}

/**
 * @createdby: Syed Muhammad Ammar
 * @description: This function is used to change the block, It doest not change the hashes of the blocks.
 * @param: instance of blockchain, and  reader *bufio.Reader for reading the input from the user.
 **/

func changeBlock(bc *MidLevelBlockchain.Blockchain, reader *bufio.Reader) {
	if len(bc.Blocks) == 0 {
		fmt.Println("No blocks to change.")
		return
	}

	fmt.Print("Enter the index of the block to change: ")
	indexStr, _ := reader.ReadString('\n')
	indexStr = strings.TrimSpace(indexStr)
	index, err := strconv.Atoi(indexStr)
	if err != nil || index < 0 || index >= len(bc.Blocks) {
		fmt.Println("Invalid block index.")
		return
	}

	fmt.Print("Enter the new transaction: ")
	newTransaction, _ := reader.ReadString('\n')
	newTransaction = strings.TrimSpace(newTransaction)

	if len(newTransaction) == 0 {
		fmt.Println("Transaction cannot be empty.")
		return
	}

	bc.ChangeBlock(index, newTransaction)
	fmt.Println("Block updated successfully.")
}

/**
 * @createdby: Syed Muhammad Ammar
 * @description: This function is used to verify the blockchain
 * @param: instance of blockchain
 **/

func verifyBlockchain(bc *MidLevelBlockchain.Blockchain) {
	isValid := bc.VerifyChain()
	if isValid {
		fmt.Println("Blockchain is valid.")
	} else {
		fmt.Println("Blockchain is invalid.")
	}
}

/**
 * @createdby: Syed Muhammad Ammar
 * @description: This function is used to set the number of transactions per block
 * @param: instance of blockchain, and  reader *bufio.Reader for reading the input from the user.
 **/

func setNumberOfTransactionsPerBlock(bc *MidLevelBlockchain.Blockchain, reader *bufio.Reader) {
	fmt.Print("Enter number of transactions per block: ")
	numStr, _ := reader.ReadString('\n')
	numStr = strings.TrimSpace(numStr)
	num, err := strconv.Atoi(numStr)
	if err != nil || num <= 0 {
		fmt.Println("Invalid number. Please enter a positive integer.")
		return
	}

	bc.SetNumberOfTransactionsPerBlock(num)
	fmt.Printf("Number of transactions per block set to %d.\n", num)
}
