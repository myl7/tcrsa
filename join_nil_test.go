package tcrsa_test

import (
	"crypto/sha256"
	"encoding/base64"
	"math/big"
	"testing"

	"github.com/myl7/tcrsa"
)

func TestJoinNilMetadataReturnsError(t *testing.T) {
	if _, err := (tcrsa.SigShareList{}).Join([]byte{1}, nil); err == nil {
		t.Fatal("Join accepted nil key metadata")
	}
}

func BenchmarkJoinExistingFixedKey(b *testing.B) {
	decode := func(s string) *big.Int {
		data, err := base64.StdEncoding.DecodeString(s)
		if err != nil {
			b.Fatal(err)
		}
		return new(big.Int).SetBytes(data)
	}
	shares, meta, err := tcrsa.NewKey(keyTestFixedSize, keyTestK, keyTestL, &tcrsa.KeyMetaArgs{
		P: decode(keyTestFixedP), Q: decode(keyTestFixedQ), R: decode(keyTestFixedR), U: decode(keyTestFixedU),
	})
	if err != nil {
		b.Fatal(err)
	}
	hash := sha256.Sum256([]byte(keyTestMessage))
	doc, err := tcrsa.PrepareDocumentHash(meta.PublicKey.Size(), keyTestHashType, hash[:])
	if err != nil {
		b.Fatal(err)
	}
	sigs := make(tcrsa.SigShareList, keyTestK)
	for i := range sigs {
		sigs[i], err = shares[i].Sign(doc, keyTestHashType, meta)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := sigs.Join(doc, meta); err != nil {
			b.Fatal(err)
		}
	}
}
