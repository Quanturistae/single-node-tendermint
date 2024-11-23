package consensus

import (
	"log"

	"github.com/Quanturistae/single-node-tendermint/crypto"
)

func StartConsensus() {
	log.Println("From consensus: Starting single-node")
    blockData := []byte("This is just a simple block message to be signed using Dilithium2")

    // Generate a key pair for signing and verification
    keys := dilithium.GenKeyPair()

    // Sign the block
    signature := dilithium.SignBlock(blockData, keys.PrivateKey)
    log.Println("From consensus: Block signed with Dilithium.")

    // Verify the block signature
    if dilithium.VerifyBlock(blockData, signature, keys.PublicKey) {
        log.Println("From consensus: Block verified successfully.")
    } else {
        log.Println("From consensus: Block verification failed.")
    }
}
