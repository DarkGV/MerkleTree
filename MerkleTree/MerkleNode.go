package MerkleTree

import (
	"crypto/sha512"
	"fmt"
)

type MerkleNode struct {
	dataHash   [32]byte
	Data       []byte
	parentNode *MerkleNode
	LeftChild  *MerkleNode
	RightChild *MerkleNode
}

func NewMerkleNode(data []byte) *MerkleNode {
	return &MerkleNode{
		dataHash:   sha512.Sum512_256(data),
		Data:       data,
		parentNode: nil,
		LeftChild:  nil,
		RightChild: nil,
	}
}

func (merkleNode *MerkleNode) AddNode(newMerkleNode *MerkleNode) *MerkleNode {
	var newRoot *MerkleNode
	fmt.Printf("%s, %v\n", newMerkleNode.Data, merkleNode.isComplete())
	if merkleNode.isComplete() {
		// Update the root
		newRoot = &MerkleNode{
			dataHash:   sha512.Sum512_256(append(merkleNode.dataHash[:], newMerkleNode.dataHash[:]...)),
			Data:       []byte{},
			parentNode: nil,
			LeftChild:  merkleNode,
			RightChild: newMerkleNode,
		}

		merkleNode.parentNode = newRoot
		newMerkleNode.parentNode = newRoot
	} else {
		merkleNode.RightChild = merkleNode.RightChild.AddNode(newMerkleNode)
		newRoot = merkleNode
	}

	return newRoot
}

// Internal Functions
func (merkleNode *MerkleNode) isComplete() bool {
	if merkleNode.count() == 0 { // this leaf is the only existing leaf
		return true
	}

	lChild := merkleNode.LeftChild.count()
	rChild := merkleNode.RightChild.count()

	if lChild == rChild {
		return true
	}
	return false
}

func (merkleNode *MerkleNode) count() int {
	if merkleNode.LeftChild == nil && merkleNode.RightChild == nil {
		return 0
	}
	return 1 + merkleNode.LeftChild.count() + merkleNode.RightChild.count()
}
