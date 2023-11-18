package MidLevelBlockchain

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

// Blockchain represents a blockchain.
type Blockchain struct {
	Blocks []*Block
}

/**
 * @createdby: Syed Muhammad Ammar
 * @description: This function is used to MineBlock mines a new block for the given transaction and previous hash
 * @description: It also adds the block to the local copy of blockchain
 * @consensus: The PoW mechanism will ensure that the hash of the block starts with a certain number of zeros.
 * @param: transactions in string , previousHash string
 * @return: instance of block
 **/

const initialDifficulty = 2            // Initial difficulty level
const difficultyAdjustmentInterval = 5 // Interval at which the difficulty will increase

func (bc *Blockchain) MineBlock(transactions []string, previousHash string) *Block {
	var nonce int = 0

	// Determine the current difficulty
	currentDifficulty := initialDifficulty + (len(bc.Blocks) / difficultyAdjustmentInterval)

	// Check if enough transactions are available to mine a block
	if len(transactions) < minTransactionsPerBlock {
		fmt.Println("Not enough transactions to mine a new block.")
		return nil
	}

	fmt.Println("Mining a new block")
	block := NewBlock(transactions, nonce, previousHash)

	for {
		hash := block.CalculateHash()
		if isValidHash(hash, currentDifficulty) {
			fmt.Printf("Block mined with hash: %s\n", hash)
			block.CurrentHash = hash
			break
		} else {
			nonce++
			block.Nonce = nonce // Update the nonce for the next iteration
		}
	}

	// Add the new block to the blockchain
	bc.Blocks = append(bc.Blocks, block)
	fmt.Println("Block added successfully.")

	return block
}

/**
 * @createdby: Syed Muhammad Ammar
 * @description: This function is used to check if the hash is valid or not
 * @param: hash string , difficulty int
 * @return: bool
 **/

func isValidHash(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hash, prefix)
}

/**
 * @createdby: Syed Muhammad Ammar
 * @description: This function is used to display the blocks in a tabular format
 * @param: instance of blockchain
 **/

func (bc *Blockchain) DisplayBlocks() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintln(w, "Block\tTransaction\tNonce\tPrevious Hash\tCurrent Hash")
	for i, block := range bc.Blocks {
		// Limit hash display to 16 characters and append "..." if it exceeds that length
		prevHash := limitHashDisplay(block.PreviousHash, 16)
		currHash := limitHashDisplay(block.CurrentHash, 16)

		fmt.Fprintf(w, "%d\t%s\t%d\t%s\t%s\n", i, strings.Join(block.Transactions, ", "), block.Nonce, prevHash, currHash)

	}

	w.Flush()
}

/**
 * @createdby: Syed Muhammad Ammar
 * @description: This function is used to change output format, It doest not change the hashes of the blocks.
 * @limitHashDisplay: limits the hash to a specified length and appends "..." if it exceeds that length.
 * @param: hash string , maxLength int
 * @return: string
 **/

func limitHashDisplay(hash string, maxLength int) string {
	if len(hash) > maxLength {
		return hash[:maxLength-3] + "..."
	}
	return hash
}

/**
 * @createdby: Syed Muhammad Ammar
 * @description: This function is used to change the block, It doest not change the hashes of the blocks.
 * @param: instance of blockchain, and  reader *bufio.Reader for reading the input from the user.
 **/

func (bc *Blockchain) ChangeBlock(blockIndex int, newTransaction string) {
	if blockIndex < 0 || blockIndex >= len(bc.Blocks) {
		fmt.Println("Invalid block index")
		return
	}

	for i := blockIndex; i < len(bc.Blocks); i++ {
		if i > 0 {
			bc.Blocks[i].PreviousHash = bc.Blocks[i-1].CurrentHash
		}

		if i == blockIndex {
			bc.Blocks[i].Transactions = append(bc.Blocks[i].Transactions, newTransaction)
		}

		bc.Blocks[i].CurrentHash = bc.Blocks[i].CalculateHash()
	}
}

/**
 * @createdby: Syed Muhammad Ammar
 * @description: This function is used to verify the chain
 * @param: instance of blockchain
 * @return: bool
 **/

func (bc *Blockchain) VerifyChain() bool {
	for i := 0; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]

		// Check block's content hash matches its current hash
		if currentBlock.CurrentHash != currentBlock.CalculateHash() {
			return false
		}

		// For all blocks except the first, check if previous hash matches
		if i > 0 {
			previousBlock := bc.Blocks[i-1]
			if currentBlock.PreviousHash != previousBlock.CurrentHash {
				return false
			}
		}

		// Check Merkle root integrity
		var txData [][]byte
		for _, tx := range currentBlock.Transactions {
			txData = append(txData, []byte(tx))
		}

		tree := NewMerkleTree(txData)
		if currentBlock.MerkleRoot != hex.EncodeToString(tree.RootNode.Data) {
			return false
		}
	}
	return true
}

/**
 * @createdby: Syed Muhammad Ammar
 * @description: This function is used to set the number of transactions per block
 * @param: instance of blockchain, and  reader *bufio.Reader for reading the input from the user.
 * @Assumption:  we add a variable to keep track of the number of transactions per block
 **/

var minTransactionsPerBlock int = 2 // default
func (bc *Blockchain) SetNumberOfTransactionsPerBlock(num int) {
	minTransactionsPerBlock = num
}
