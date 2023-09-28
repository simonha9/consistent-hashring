package kvstore

import (
	"hash/fnv"

	"github.com/simonha9/consistent-hashring/pkg"
)

// This is an implementation of a key value store using consistent hashring as the underlying data structure
// KVStore is a struct that holds the state of the KVStore

type KVStore struct {
	HashRing *pkg.ConsistentHashRing
}

// NewKVStore creates a new KVStore, this needs to be uint32 not byte
func NewKVStore() *KVStore {

	k := &KVStore{
		HashRing: pkg.NewConsistentHashRing(fnv.New64(), []uint32{}),
	}

	nodes := []*pkg.Node{
		{
			Key:          "server1",
			HashedKey:    0,
			Contents:     []pkg.Key{},
			ParentServer: pkg.Server{ServerName: "server1"},
		},
		{
			Key:          "server2",
			HashedKey:    0,
			Contents:     []pkg.Key{},
			ParentServer: pkg.Server{ServerName: "server2"},
		},
	}

	for _, node := range nodes {
		k.HashRing.AddServer(node)
	}

	return k
}

func (kv *KVStore) Get(key string) (string, error) {
	return "", nil
}

func (kv *KVStore) Set(key string, value string) error {
	// hash the key and value onto the hashring
	return kv.HashRing.AddKey(key)
}
