package pkg

type Key struct {
	key       string // Key of the node
	value     string
	hashedKey uint64 // hashed key on the hashring
}

func (k *Key) GetKey() string {
	return k.key
}

func (k *Key) GetValue() string {
	return k.value
}

func (k *Key) GetHashedKey() uint64 {
	return k.hashedKey
}
