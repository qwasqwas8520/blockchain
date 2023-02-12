package blockchain

import (
	"bytes"
	"crypto/sha256"
	"golang/utils"
	"time"
)

type Block struct {
	Timestamp int64
	Hash      []byte
	PrevHash  []byte
	Data      []byte
	Nonce     int64
	Target    []byte
}

func (block *Block) SetHash() {
	information := bytes.Join([][]byte{utils.ToHexInt(block.Timestamp), block.PrevHash, block.Data, utils.ToHexInt(block.Nonce), block.Target}, []byte{})
	hash := sha256.Sum256(information)
	block.Hash = hash[:]
}

func CreateBlock(prevhash, data []byte) *Block {
	block := Block{time.Now().Unix(), []byte{}, prevhash, data, 0, []byte{}}
	block.Target = block.GetTarget()
	block.Nonce = block.FindNonce()
	block.SetHash()
	return &block
}

func GenesisBlock() *Block {
	genesisWords := "創世區塊"
	return CreateBlock([]byte{}, []byte(genesisWords))
}
