package pkg

import "fmt"

// ConsistentHashRing is a struct that holds the state of the consistent hash ring
type ConsistentHashRing struct {
	hashFunction func(string) uint32
	Nodes        map[uint32]*Node // map of hashed keys to nodes
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
		Nodes:        make(map[uint32]*Node),
	}
	// quicksort keys
	quickSort(keys, 0, len(keys)-1)

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
	server := cr.findClosestServer(h)
	server.contents = append(server.contents, value)
}

func (cr *ConsistentHashRing) findClosestServer(hash uint32) *Node {
	// find the closest server to the hash
	// use binary search to find the closest server (leetcode)

	low := 0
	high := len(cr.Keys) - 1
	for low < high {
		mid := (low + high) / 2
		if cr.Keys[mid].hashedKey == hash {
			return cr.Nodes[hash]
		} else if cr.Keys[mid].hashedKey < hash {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return cr.Nodes[cr.Keys[low].hashedKey]
}
