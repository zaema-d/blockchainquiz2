// blockchain_simple.go
package main

import (
    "fmt"
    "strconv"
    "time"
)

// Define the structure of each block in the blockchain
type SimpleBlock struct {
    TimeCreated   string
    Information   string
    PrevBlockHash string
}

// Define the blockchain as a slice of blocks
type SimpleBlockchain struct {
    simpleBlocks []SimpleBlock
}

// Function to create a new block given some information and the hash of the previous block
func CreateNewBlock(info string, prevHash string) SimpleBlock {
    return SimpleBlock{TimeCreated: time.Now().Format(time.RFC3339), Information: info, PrevBlockHash: prevHash}
}

// Function to add a new block to the blockchain
func (bc *SimpleBlockchain) AddSimpleBlock(info string) {
    var prevHash string
    if len(bc.simpleBlocks) > 0 {
        prevHash = bc.simpleBlocks[len(bc.simpleBlocks)-1].PrevBlockHash
    }
    newBlock := CreateNewBlock(info, prevHash)
    bc.simpleBlocks = append(bc.simpleBlocks, newBlock)
}

// Function to initialize the blockchain with a genesis block
func InitializeSimpleBlockchain() *SimpleBlockchain {
    genesisBlock := CreateNewBlock("Genesis Block", "")
    return &SimpleBlockchain{simpleBlocks: []SimpleBlock{genesisBlock}}
}

// Function to display the contents of the blockchain
func (bc *SimpleBlockchain) DisplaySimpleBlockchain() {
    for i, block := range bc.simpleBlocks {
        fmt.Printf("Block #%d\n", i)
        fmt.Printf("Created: %s\n", block.TimeCreated)
        fmt.Printf("Information: %s\n", block.Information)
        fmt.Printf("Prev. Hash: %s\n\n", block.PrevBlockHash)
    }
}

func main() {
    // Initialize the blockchain with a genesis block
    myBlockchain := InitializeSimpleBlockchain()
    
    // Add some blocks with dummy information
    myBlockchain.AddSimpleBlock("First Block after Genesis")
    myBlockchain.AddSimpleBlock("Second Block after Genesis")
    myBlockchain.AddSimpleBlock("Third Block after Genesis")
    
    // Display the blockchain
    myBlockchain.DisplaySimpleBlockchain()
}
