# MerkleTree
A simple MerkleTree implementation in Go

# Usage
It is very simple to use this library

```go
package main

import "./MerkleTree"

func main(){
    merkleRoot := MerkleTree.NewMerkleNode([]byte("Example"))

    merkleRoot = merkleRoot.AddNode(MerkleTree.NewMerkleNode([]byte("Example1")))
}
```