package gochain

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// Block represent a record in a blockchain
type Block struct {
	Index         uint64      `json:"index"`         // Index is the index of a block
	PreviousBlock *Block      `json:"previousBlock"` // PreviousBlock points to the previous block
	Timestamp     int64       `json:"timestamp"`     // Timestamp is the time of the record creation
	Data          interface{} `json:"data"`          // Data is the value of a record
	Hash          string      `json:"hash"`          // Hash is hash identifier for a block
}

// NewBlock returns a new empty block
func NewBlock() *Block {
	return &Block{
		Index:         0,
		PreviousBlock: nil,
		Timestamp:     time.Now().Unix(),
		Data:          []byte(""),
		Hash:          hex.EncodeToString(sha256.New().Sum(nil)),
	}
}

// MakeHash returns a new hash string
func MakeHash(payload string) (string, error) {
	h := sha256.New()
	_, err := h.Write([]byte(payload))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}
