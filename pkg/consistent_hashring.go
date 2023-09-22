package pkg

import "fmt"

// Consistent hash ring has keys and servers, you hash the node
// then store keys on the servers.
// the server that a key belongs to is, hash the key then go clockwise to find
// the closest server and that server holds the key

// ConsistentHashRing is a struct that holds the state of the consistent hash ring
type ConsistentHashRing struct {
	hashFunction func(string) uint32
	Nodes        []*Node
	Keys         []*Key
}

type Node struct {
	key       string // Key of the node
	hashedKey uint32 // hashed key to identify which key it belongs to
	contents  []Key  // these are servers so we need to hold information
}

type Key struct {
	key       string // Key of the node
	hashedKey uint32 // hashed key on the hashring
}

// NewConsistentHashRing creates a new ConsistentHashRing
func NewConsistentHashRing(hash func(string) uint32, keys []uint32) *ConsistentHashRing {
	cr := &ConsistentHashRing{
		hashFunction: hash,
		Nodes:        make([]*Node, 0),
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
func (cr *ConsistentHashRing) AddKey(key string) {
	h := cr.hashFunction(key)
	k := Key{
		key:       key,
		hashedKey: h,
	}

	// find the server that is closest to the value
	server := cr.findClosestServer(h)
	server.contents = append(server.contents, k)
}

func (cr *ConsistentHashRing) findClosestServer(hash uint32) *Node {
	// find the closest server to the hash
	// use binary search to find the closest server (leetcode)

	low := 0
	high := len(cr.Nodes) - 1
	for low < high {
		mid := (low + high) / 2
		if cr.Nodes[mid].hashedKey == hash {
			return cr.Nodes[hash]
		} else if cr.Nodes[mid].hashedKey < hash {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return cr.Nodes[low]
}

// Now to support multiple adding and removing / redistribution of keys

func (cr *ConsistentHashRing) AddNode(node *Node) {
	// add the node to the hash ring
	// redistribute the keys

	// find all the keys between this node and the next node to redistribute them, by
	// finding all the keys < newly added node
	// that is, find the largest key such that k < newly added node
	// and every key between k and j where j is the smallest key > previous node inclusive

}

func (cr *ConsistentHashRing) RemoveNode(node *Node) {
	// remove the node from the hash ring
	// redistribute the keys

	// find all the keys between this node and
}

/*
* But theres not guarantee that the keys are evenly distributed
* So we need to add virtual nodes to the hash ring, TODO: implement virtual nodes
 */
