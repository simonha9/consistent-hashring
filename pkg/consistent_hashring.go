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
	Keys         *BSTNode
}

// NewConsistentHashRing creates a new ConsistentHashRing
func NewConsistentHashRing(hash func(string) uint32, keys []uint32) *ConsistentHashRing {
	cr := &ConsistentHashRing{
		hashFunction: hash,
		Nodes:        NewBSTNode(&Node{"", 0, nil, Server{"serverroot"}}),
	}
	// quicksort keys
	cr.Keys = NewBSTNode(&Node{"", 0, nil, Server{"keyroot"}})

	nodes := []*Node{}
	bKeys := []*Key{}

	for _, key := range keys {
		h := hash(fmt.Sprintf("%d", key))

		k := &Key{
			key:       fmt.Sprintf("%d", key),
			hashedKey: h,
		}
		n := &Node{
			key:       fmt.Sprintf("%d", key),
			hashedKey: h,
			contents:  []Key{*k},
		}

		nodes = append(nodes, n)
		bKeys = append(bKeys, k)
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

	// find the server that is closest to the value should be right biased
	server := cr.findNode(h, false, false)
	server.contents = append(server.contents, k)
}

func (cr *ConsistentHashRing) findNode(hash uint32, left bool, exact bool) *Node {
	// left biased binsearch but just search through bstnode interface
	// find the node that is closest to the hash value

	// find the node
	if left {
		return cr.Nodes.SearchLeftBiased(hash, exact)
	}
	return cr.Nodes.SearchRightBiased(hash, exact)
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

	leftSibling := cr.findNode(node.hashedKey, true, false)
	nextLargestKey := cr.findNode(node.hashedKey, false, false)

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
	cr.addServer(node)
}

func (cr *ConsistentHashRing) addServer(node *Node) {
	cr.Nodes.Insert(node)
}

func (cr *ConsistentHashRing) RemoveServer(node *Node) {
	// remove the node from the hash ring
	// redistribute the keys

	// find all the keys between this node and last one
	// all the keys here need to be redistributed to the next sibling

	rightSibling := cr.findNode(node.hashedKey, false, false)
	copy(rightSibling.contents, append(rightSibling.contents, node.contents...))
	cr.Nodes.Delete(node.hashedKey)
}

/*
But theres no guarantee that the keys are evenly distributed
So we need to add virtual nodes to the hash ring, TODO: implement virtual nodes

But this is not a change in BST structure its just a change in number of server keys
I guess rather a node will have multiple servers and each server will hold keys but how
do we actually do this

so BSTNode will have a list of Nodes of the same server, or we can hide it in the "node" struct
with a list of servers, but this might not be the best because we need to find what server a key belongs to
since they will be in different places in the ring

so each virtualnode is a node, and each node will belong to a server
so then we don't really care the servers but rather they are all just different

Then maybe we have a Server struct that holds a list of nodes in case we need to do any checks like
how many there are etc etc

*/
