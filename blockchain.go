package gochain

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// ErrFirstBlockMustBeGenesisBlock indicates that the first block of a blockchain must be created with NewBlockChaim
var ErrFirstBlockMustBeGenesisBlock = errors.New("The first block must be created in NewBlockChain")

// BlockChain groups all Blocks in a chain
type BlockChain struct {
	sync.Mutex `json:"-"`
	Blocks     []*Block `json:"blocks"`
}

// NewBlockChain return a new blockchain
func NewBlockChain(genesisBlock *Block) *BlockChain {
	return &BlockChain{Blocks: []*Block{genesisBlock}}
}

// NextBlock generate the next block in a blockchain
func (b *BlockChain) NextBlock(blockData interface{}) (*Block, error) {
	if len(b.Blocks) == 0 {
		return nil, ErrFirstBlockMustBeGenesisBlock
	}

	b.Lock()
	defer b.Unlock()

	previousBlock := b.GetLatestBlock()

	block := NewBlock()
	block.Index = previousBlock.Index + 1
	block.Timestamp = time.Now().Unix()
	block.PreviousBlock = previousBlock
	block.Data = blockData

	hash, err := MakeHash(fmt.Sprintf("%d%s%d", block.Index, previousBlock.Hash, block.Timestamp))
	if err != nil {
		return nil, err
	}
	block.Hash = hash

	b.Blocks = append(b.Blocks, block)

	return block, nil
}

// GetLatestBlock retrieves the latest block from the blockchain
func (b *BlockChain) GetLatestBlock() *Block {
	if len(b.Blocks) == 0 {
		return nil
	}

	return b.Blocks[len(b.Blocks)-1]
}
