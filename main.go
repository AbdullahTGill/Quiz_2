package main

import (
	"fmt"
)

// Block represents a data block
type Block struct {
	ID   int
	Data string
}

// BlockChain represents a collection of blocks
type BlockChain struct {
	Blocks []Block
}

// DisplayAllBlocks prints all blocks in the blockchain
func (bc *BlockChain) DisplayAllBlocks() {
	fmt.Println("Blocks in the Blockchain:")
	for _, block := range bc.Blocks {
		fmt.Printf("Block ID: %d, Data: %s\n", block.ID, block.Data)
	}
}

// NewBlock creates a new block and adds it to the blockchain
func (bc *BlockChain) NewBlock(id int, data string) {
	newBlock := Block{ID: id, Data: data}
	bc.Blocks = append(bc.Blocks, newBlock)
	fmt.Println("New block added to the Blockchain.")
}

// ModifyBlock modifies the data of a block with the given ID
func (bc *BlockChain) ModifyBlock(id int, newData string) {
	for i := range bc.Blocks {
		if bc.Blocks[i].ID == id {
			bc.Blocks[i].Data = newData
			fmt.Printf("Block with ID %d modified.\n", id)
			return
		}
	}
	fmt.Printf("Block with ID %d not found.\n", id)
}

func main() {
	// Example usage
	bc := &BlockChain{}
	bc.NewBlock(1, "First Block")
	bc.NewBlock(2, "Second Block")
	bc.DisplayAllBlocks()

	// Modify a block
	bc.ModifyBlock(2, "Updated Second Block")
	bc.DisplayAllBlocks()
}
