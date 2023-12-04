package pkg

import (
	"bytes"
	"crypto/sha256"
)

// MerkleNode represents a node in the Merkle tree.
type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

// MerkleTree represents a Merkle tree.
type MerkleTree struct {
	RootChecksum []byte
}

// GetMerkleRoot returns the root hash of the Merkle tree.
func (mt *MerkleTree) GetMerkleRootChecksum() []byte {
	return mt.RootChecksum
}

// NewMerkleTree constructs a Merkle tree from transaction data.
func NewMerkleTree(Transactions []*Transaction) *MerkleTree {
	var nodes []MerkleNode

	// Create leaf nodes from transactions data
	for _, Tx := range Transactions {
		node := MerkleNode{Data: Tx.Data}
		nodes = append(nodes, node)
	}

	// Build the Merkle tree by calculating root value
	tree := buildTree(nodes)

	return tree
}

// buildTree constructs the Merkle tree using a bottom-up approach.
func buildTree(nodes []MerkleNode) *MerkleTree {

	for len(nodes) > 1 {
		var level []MerkleNode

		// If the number of nodes is odd, replicate the last node
		if len(nodes)%2 != 0 {
			lastNode := nodes[len(nodes)-1]
			nodes = append(nodes, lastNode)
		}

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
			level = append(level, node)
		}
		// Move to the next level in the tree
		nodes = level
	}

	return &MerkleTree{RootChecksum: nodes[0].Data}
}

// concatenateAndHash combines two byte slices and computes their SHA-256 hash.
func concatenateAndHash(left, right []byte) []byte {
	hash := sha256.Sum256(append(left, right...))
	return hash[:]
}

func ValidateBlockIntegrity(b *Block) bool {
	merkleRootChecksum := NewMerkleTree(b.Transactions).GetMerkleRootChecksum()

	return bytes.Equal(merkleRootChecksum, b.MerkleRootChecksum)
}
