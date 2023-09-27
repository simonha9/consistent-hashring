package kvstore

import (
	"crypto/sha256"

	"github.com/simonha9/consistent-hashring/pkg"
)

// This is an implementation of a key value store using consistent hashring as the underlying data structure
// KVStore is a struct that holds the state of the KVStore

type KVStore struct {
	HashRing *pkg.ConsistentHashRing
}

// NewKVStore creates a new KVStore, this needs to be uint32 not byte
func NewKVStore() *KVStore {
	hash := sha256.New224().Sum32()
	return &KVStore{
		HashRing: pkg.NewConsistentHashRing(sha256.New(), []uint32{}),
	}
}
