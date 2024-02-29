package main

import (
    "fmt"
    "strconv"
    "time"
)

type BlockStructure struct {
    CreationTime string
    Payload      string
    PreviousHash string
    BlockID      int
}

type MiniBlockchain struct {
    blocks []BlockStructure
}

// SimpleHash simulates a hashing function for educational purposes.
func SimpleHash(input string) string {
    // This is a very basic "hashing" mechanism for demonstration.
    // In real applications, you should use a cryptographic hash function.
    return fmt.Sprintf("%x", len(input))
}

func GenerateBlock(payload string, previousHash string, blockID int) BlockStructure {
    block := BlockStructure{
        CreationTime: time.Now().Format("2006-01-02 15:04:05"),
        Payload:      payload,
        PreviousHash: previousHash,
        BlockID:      blockID,
    }
    // Generate a simple hash for the PreviousHash field based on Payload
    block.PreviousHash = SimpleHash(block.Payload + strconv.Itoa(block.BlockID))
    return block
}

func (mbc *MiniBlockchain) AppendBlock(payload string) {
    var previousHash string
    var blockID int
    if len(mbc.blocks) > 0 {
        previousBlock := mbc.blocks[len(mbc.blocks)-1]
        previousHash = previousBlock.PreviousHash
        blockID = previousBlock.BlockID + 1
    } else {
        previousHash = "0"
        blockID = 1
    }
    newBlock := GenerateBlock(payload, previousHash, blockID)
    mbc.blocks = append(mbc.blocks, newBlock)
}

func InitBlockchain() *MiniBlockchain {
    genesisBlock := GenerateBlock("Start of Chain", "0", 0)
    return &MiniBlockchain{blocks: []BlockStructure{genesisBlock}}
}

func (mbc *MiniBlockchain) ShowBlockchain() {
    for _, block := range mbc.blocks {
        fmt.Println("---------------")
        fmt.Printf("Block ID: %d\n", block.BlockID)
        fmt.Printf("Creation Time: %s\n", block.CreationTime)
        fmt.Printf("Payload: %s\n", block.Payload)
        fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
        fmt.Println("---------------\n")
    }
}

func main() {
    bc := InitBlockchain()

    bc.AppendBlock("First real transaction")
    bc.AppendBlock("Second real transaction")
    bc.AppendBlock("Third real transaction")

    bc.ShowBlockchain()
}
