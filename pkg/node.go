package pkg

import "fmt"

type Node struct {
	Key          string // Key of the node
	HashedKey    uint64 // hashed key to identify which key it belongs to
	Contents     []Key  // these are servers so we need to hold information
	ParentServer Server
}

func (n *Node) GetKey(key string) (string, error) {
	for _, k := range n.Contents {
		if k.key == key {
			return k.GetValue(), nil
		}
	}
	return "", fmt.Errorf("key not found")
}
