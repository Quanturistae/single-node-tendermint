package main

import (
	"log"
    "os"

    "github.com/Quanturistae/single-node-tendermint/consensus"
)

func main() {
    log.Println("From main: Starting Tendermint single-node network with Dilithium...")
    
    go consensus.StartConsensus()
    go rpc.StartRPCServer()

    // Prevent the main function from exiting
    //select {}
}