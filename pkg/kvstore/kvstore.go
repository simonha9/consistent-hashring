package kvstore

import (
	"github.com/segmentio/fasthash/fnv1a"
	"github.com/simonha9/consistent-hashring/pkg"
)

// This is an implementation of a key value store using consistent hashring as the underlying data structure
// KVStore is a struct that holds the state of the KVStore

type KVStore struct {
	HashRing *pkg.ConsistentHashRing
}

// NewKVStore creates a new KVStore, this needs to be uint32 not byte
func NewKVStore() *KVStore {

	return &KVStore{
		HashRing: pkg.NewConsistentHashRing(, []uint32{}),
	}
}
