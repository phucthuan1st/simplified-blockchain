package pkg

import (
	"crypto/sha256"
	"encoding/hex"
)

// MerkleNode represents a node in the Merkle tree.
type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

// MerkleTree represents a Merkle tree.
type MerkleTree struct {
	Root *MerkleNode
}

// NewMerkleTree constructs a Merkle tree from transaction data.
func NewMerkleTree(transactions [][]byte) *MerkleTree {
	var nodes []MerkleNode
	// Create leaf nodes from transaction data
	for _, txData := range transactions {
		node := MerkleNode{Data: txData}
		nodes = append(nodes, node)
	}
	// Build the Merkle tree
	root := buildTree(nodes)
	return &MerkleTree{Root: &root[0]}
}

// buildTree constructs the Merkle tree using a bottom-up approach.
func buildTree(nodes []MerkleNode) []MerkleNode {
	// If the number of nodes is odd, replicate the last node
	if len(nodes)%2 != 0 {
		lastNode := nodes[len(nodes)-1]
		nodes = append(nodes, lastNode)
	}
	var newLevel []MerkleNode
	for len(nodes) > 1 {
		newLevel = nil // Clear newLevel slice

		// Combine nodes in pairs to create higher-level nodes
		for i := 0; i < len(nodes); i += 2 {
			node := MerkleNode{
				Left:  &nodes[i],
				Right: nil,
				Data:  concatenateAndHash(nodes[i].Data, nodes[i+1].Data),
			}
			if i+1 < len(nodes) {
				node.Right = &nodes[i+1]
			}
			newLevel = append(newLevel, node)
		}
		// Move to the next level in the tree
		nodes = newLevel
	}
	return nodes
}

// concatenateAndHash combines two byte slices and computes their SHA-256 hash.
func concatenateAndHash(left, right []byte) []byte {
	hash := sha256.Sum256(append(left, right...))
	return hash[:]
}

// GetMerkleRoot returns the root hash of the Merkle tree.
func (mt *MerkleTree) GetMerkleRoot() string {
	return hex.EncodeToString(mt.Root.Data)
}

// GetMerkleRootData returns the root node's data of the Merkle tree.
func (mt *MerkleTree) GetMerkleRootData() []byte {
	return mt.Root.Data
}
