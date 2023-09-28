package pkg

type Node struct {
	Key          string // Key of the node
	HashedKey    uint64 // hashed key to identify which key it belongs to
	Contents     []Key  // these are servers so we need to hold information
	ParentServer Server
}
