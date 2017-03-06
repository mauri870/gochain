package gochain

import "testing"

func TestNextBlock(t *testing.T) {
	howManyBlocks := []int{1, 3, 10}

	block := NewBlock()

	for _, hmb := range howManyBlocks {
		blockchain := NewBlockChain(block)
		for i := hmb; i > 0; i-- {
			tmpBlock, err := blockchain.NextBlock(i)
			if err != nil {
				t.Error(err)
			}

			if tmpBlock.Data != i {
				t.Errorf("Block data must be %d, got %v", i, tmpBlock.Data)
			}
		}

		if len(blockchain.Blocks) != hmb+1 {
			t.Errorf("The blockchain length must be %d, got %d", hmb, len(blockchain.Blocks))
		}
	}
}

func TestGetLatestBlock(t *testing.T) {
	genesisBlock := NewBlock()
	blockchain := NewBlockChain(genesisBlock)
	if blockchain.GetLatestBlock() != genesisBlock {
		t.Error(ErrFirstBlockMustBeGenesisBlock.Error())
	}

	block2, err := blockchain.NextBlock(`\o/`)
	if err != nil {
		t.Error(err)
	}

	if blockchain.GetLatestBlock() != block2 {
		t.Error("A blockchain must return the latest block if it is not empty")
	}
}
