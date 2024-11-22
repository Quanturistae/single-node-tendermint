package dilithium

import (
    "errors"
    "log"

    "github.com/open-quantum-safe/liboqs-go/oqs"
)

type keyPair struct {
    PublicKey []byte
    PrivateKey []byte
}

func GenKeyPair() (keyPair, error) {
	signer := oqs.Signature{}
	defer signer.Clean()

	err := signer.Init("Dilithium2", nil)
	if err != nil {
		log.Fatalf("Failed to initialize Dilithium signer: %v", err)
	}

	pubKey, privKey, err := signer.GenerateKeyPair()

	return keyPair{PublicKey: pubKey, PrivateKey: privKey}, nil
}

func SignBlock(message []byte, privKey []byte) ([]byte, error) {
	signer := oqs.Signature{}
	defer signer.Clean()

	err := signer.Init("Dilihium2", nil)
	if err != nil {
		return nil, err
	}

	signature, err := signer.Sign(message, privKey)
	if err != nil {
		return nil, err
	}

	return signature, nil
}


func VerifyBlock(message []byte, signature []byte, pubKey []byte) bool {
	verifier := oqs.Signature{}
	defer verifier.Clean()

	err := verifier.Init("Dilithium2", nil)
	if err != nil {
		return false
	}

	valid, err := verifier.Verify(message, signature, pubKey)
	if err != nil {
		return false
	}

	return valid
}
