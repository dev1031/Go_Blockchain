package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

type BlockChain struct {
	blocks []*Block
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevHash := chain.blocks[len(chain.blocks)-1].Hash
	new := CreateBlock(data, prevHash)
	chain.blocks = append(chain.blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	//fmt.Println(&BlockChain{[]*Block{Genesis()}})
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()
	chain.AddBlock("First block after genesis")
	chain.AddBlock("Second block after genesis")
	chain.AddBlock("Third block after genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Prev hash : %x\n", block.PrevHash)
		fmt.Printf("Data in block : %s\n", block.Data)
		fmt.Printf("Hash : %x\n", block.Hash)
	}
}
