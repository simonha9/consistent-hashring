package pkg

type Node struct {
	key          string // Key of the node
	hashedKey    uint64 // hashed key to identify which key it belongs to
	contents     []Key  // these are servers so we need to hold information
	parentServer Server
}
