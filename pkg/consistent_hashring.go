package pkg

import "fmt"

// ConsistentHashRing is a struct that holds the state of the consistent hash ring
type ConsistentHashRing struct {
	hashFunction func(string) uint32
	Nodes        []*Node
	Keys         []*Key
}

type Node struct {
	key       string   // Key of the node
	hashedKey uint32   // hashed key to identify which Key it belongs to
	contents  []string // these are servers so we need to hold information
}

type Key struct {
	key       string // Key of the node
	hashedKey uint32 // hashed key on the hashring
}

// NewConsistentHashRing creates a new ConsistentHashRing
func NewConsistentHashRing(hash func(string) uint32, keys []uint32) *ConsistentHashRing {
	cr := &ConsistentHashRing{
		hashFunction: hash,
		Nodes:        []*Node{},
	}
	for _, key := range keys {
		h := hash(fmt.Sprintf("%d", key))
		k := &Key{
			key:       fmt.Sprintf("%d", key),
			hashedKey: h,
		}
		cr.Keys[key] = k
	}
	return cr
}

// AddNode adds a node to the consistent hash ring
func (cr *ConsistentHashRing) AddValue(value string) {
	h := cr.hashFunction(value)
	// find the server that is closest to the value
}
