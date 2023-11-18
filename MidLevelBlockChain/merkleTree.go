package MidLevelBlockchain

import (
	"crypto/sha256"
	"fmt"
	"os"
	"text/tabwriter"
)

type MerkleTree struct {
	RootNode *MerkleNode
}

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

/**
 * @createdby: Syed Muhammad Ammar
 * @description: This function is used to limit the display the merkle tree hashes
 * @param: instance of merkle tree, length int for the length of the hash
 **/

func limitHash(hash []byte, length int) string {
	if len(hash) > length {
		return fmt.Sprintf("%x...", hash[:length])
	}
	return fmt.Sprintf("%x", hash)
}

/**
 * @createdby: Syed Muhammad Ammar
 * @description: This function is used to display the merkle node in a tabular format
 * @param: instance of merkle tree, level int for the level of the merkle tree
 **/

func (node *MerkleNode) displayMerkleNode(w *tabwriter.Writer, level int) {
	if node == nil {
		return
	}

	leftHash := "Leaf Node"
	rightHash := "Lead Node"

	if node.Left != nil {
		leftHash = limitHash(node.Left.Data, 16)
		node.Left.displayMerkleNode(w, level+1)
	}

	if node.Right != nil {
		rightHash = limitHash(node.Right.Data, 16)
		node.Right.displayMerkleNode(w, level+1)
	}

	resultHash := limitHash(node.Data, 16)
	fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", level, leftHash, rightHash, resultHash)
}

/**
 * @createdby: Syed Muhammad Ammar
 * @description: This function is used to display the merkle tree
 * @param: instance of merkle tree
 **/

func (tree *MerkleTree) DisplayMerkleTree() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "Level\tLeft Child\tRight Child\tResulting Hash")
	tree.RootNode.displayMerkleNode(w, 1)
	w.Flush()
}

/**
 * @createdby: Syed Muhammad Ammar
 * @description: This function is used to create the new merkle node
 * @param: left: instance of merkle node, right: instance of merkle node, data []byte: data of the merkle node in bytes
 * @return: instance of merkle node
 **/

func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	node := MerkleNode{}

	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		node.Data = hash[:]
		// fmt.Printf("Leaf Node created with data: %x\n", node.Data)
	} else {
		prevHashes := append(left.Data, right.Data...)
		hash := sha256.Sum256(prevHashes)
		node.Data = hash[:]
		// fmt.Printf("Parent Node created with left child: %x and right child: %x resulting in hash: %x\n", left.Data, right.Data, node.Data)
	}

	node.Left = left
	node.Right = right

	return &node
}

/**
 * @createdby: Syed Muhammad Ammar
 * @description: This function is used to create the new merkle tree
 * @param: data [][]byte: data of the merkle tree in bytes (transactions) in the form of 2D array of bytes
 * @return: instance of merkle tree
 **/
func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []MerkleNode

	// Create leaf nodes for each data block
	for _, datum := range data {
		node := NewMerkleNode(nil, nil, datum)
		nodes = append(nodes, *node)
	}

	//fmt.Println("Leaf nodes created...")

	// While there's more than 1 node, keep hashing till we reach root
	for len(nodes) > 1 {
		level := []MerkleNode{}

		// If there's an odd number of nodes, append the last node again
		if len(nodes)%2 != 0 {
			nodes = append(nodes, nodes[len(nodes)-1])
		}

		for i := 0; i < len(nodes); i += 2 {
			node := NewMerkleNode(&nodes[i], &nodes[i+1], nil)
			level = append(level, *node)
		}

		//fmt.Println("Next level of parent nodes created...")
		nodes = level
	}

	tree := MerkleTree{&nodes[0]}

	fmt.Println("Merkle Tree Structure:")
	tree.DisplayMerkleTree()

	return &tree
}
