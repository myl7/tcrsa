package tcrsa_test

import (
	"crypto"
	"crypto/sha256"
	"testing"

	"github.com/myl7/tcrsa"
)

const pkcs1EncodingTestMessage = "Hello World"
const pkcs1EncodingLeyLength = 64

func TestPrepareDocumentHash(t *testing.T) {
	docHash := sha256.Sum256([]byte(pkcs1EncodingTestMessage))
	docPKCS1, err := tcrsa.PrepareDocumentHash(pkcs1EncodingLeyLength, crypto.SHA256, docHash[:])
	if err != nil {
		t.Errorf("couldn't prepare document hash: %v", err)
	}
	// Check byte length
	if len(docPKCS1) != pkcs1EncodingLeyLength {
		t.Errorf("prepared hash has not the desired length")
	}

}
