package pkg

import "fmt"

// Consistent hash ring has keys and servers, you hash the node
// then store keys on the servers.
// the server that a key belongs to is, hash the key then go clockwise to find
// the closest server and that server holds the key

// ConsistentHashRing is a struct that holds the state of the consistent hash ring
type ConsistentHashRing struct {
	hashFunction func(string) uint32
	Nodes        *BSTNode
	Keys         []*Key
}

// NewConsistentHashRing creates a new ConsistentHashRing
func NewConsistentHashRing(hash func(string) uint32, keys []uint32) *ConsistentHashRing {
	cr := &ConsistentHashRing{
		hashFunction: hash,
		Nodes:        NewBSTNode(&Node{"", 0, nil}),
	}
	// quicksort keys
	quickSort(keys, 0, len(keys)-1)
	cr.Keys = make([]*Key, len(keys))

	for _, key := range keys {
		h := hash(fmt.Sprintf("%d", key))
		k := &Key{
			key:       fmt.Sprintf("%d", key),
			hashedKey: h,
		}
		cr.Nodes.Insert(&Node{
			key:       fmt.Sprintf("%d", key),
			hashedKey: h,
			contents:  []Key{*k},
		})
		cr.Keys = append(cr.Keys, k)
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
	// left biased binsearch but just search through bstnode interface
	// find the node that is closest to the hash value

	// find the node
	node := cr.Nodes.Search(hash)
	if node == nil {
		return nil
	}
	return node.Node
}

func (cr *ConsistentHashRing) findNextLargestKey(hash uint32) Key {
	// This is a right biased binsearch

	low := 0
	high := len(cr.Keys) - 1
	for low < high {
		mid := (low + high) / 2
		if hash < cr.Keys[mid].hashedKey {
			high = mid - 1
		} else {
			low = mid
		}
	}
	return *cr.Keys[low]
}

// Now to support multiple adding and removing / redistribution of keys

func (cr *ConsistentHashRing) AddServer(node *Node) {
	// Add a server to the hashring, need a redistribution of keys
	// Find the hash server value, use bin search to find placement of new server
	// but take a note of the left and right, we will need a redistribution of keys between
	// newnode-1 and curnode so that curnode will take all those keys
	// we don't need to touch any of the nodes on the right side because they are already in the right place

	// but we actually do this by finding the smallest key > cur node, then all keys in newnode-1 server
	// with values < target will get redistributed

	leftSibling := cr.findClosestServer(node.hashedKey)
	nextLargestKey := cr.findNextLargestKey(node.hashedKey)

	remove := []Key{}
	keep := []Key{}

	for _, key := range leftSibling.contents {
		if key.hashedKey < nextLargestKey.hashedKey {
			remove = append(remove, key)
		} else {
			keep = append(keep, key)
		}
	}

	copy(leftSibling.contents, keep)
	copy(node.contents, remove)
	cr.addServerAndGrow(node)
}

func (cr *ConsistentHashRing) addServerAndGrow(node *Node) {
	// TODO: needs to be perma sorted so a BST is better
	cr.Nodes.Insert(node)
}

func (cr *ConsistentHashRing) RemoveServer(node *Node) {
	// remove the node from the hash ring
	// redistribute the keys

	// find all the keys between this node and
}

/*
* But theres not guarantee that the keys are evenly distributed
* So we need to add virtual nodes to the hash ring, TODO: implement virtual nodes
 */
