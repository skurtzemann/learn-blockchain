package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
)

type BlockChain struct {
	blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (b *Block) DeriveHash() {
	// Join current block data and previous block hash
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	// Calculate hash of info
	hash := sha256.Sum256(info)
	// Then assign the result to block hash
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		[]byte{},
		[]byte(data),
		prevHash}
	block.DeriveHash()
	return block
}

// Genesis is the first block of the chain with a empty PrevHash
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// Add a block into the blockchain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

// Init a BlockChain with a first Genesis block
func InitBlockChain() *BlockChain {
	return &BlockChain{
		[]*Block{Genesis()},
	}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for i, block := range chain.blocks {
		fmt.Printf("[%s]\n", strconv.Itoa(i))
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}

}
