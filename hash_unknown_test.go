package tcrsa_test

import (
	"crypto"
	"encoding/hex"
	"strings"
	"testing"

	"github.com/myl7/tcrsa"
)

func TestPrepareDocumentHashUnknownHashReturnsError(t *testing.T) {
	if _, err := tcrsa.PrepareDocumentHash(64, crypto.Hash(255), nil); err == nil {
		t.Fatal("PrepareDocumentHash accepted an unknown hash")
	}
}

func TestPrepareDocumentHashSupportedEncodings(t *testing.T) {
	for _, tc := range []struct {
		name   string
		hash   crypto.Hash
		digest []byte
		want   string
	}{
		{"SHA256", crypto.SHA256, make([]byte, 32), "0001" + strings.Repeat("ff", 10) + "003031300d060960864801650304020105000420" + strings.Repeat("00", 32)},
		{"raw", 0, []byte{1, 2, 3}, "0001" + strings.Repeat("ff", 58) + "00010203"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tcrsa.PrepareDocumentHash(64, tc.hash, tc.digest)
			if err != nil {
				t.Fatal(err)
			}
			if hex.EncodeToString(got) != tc.want {
				t.Fatalf("encoding=%x, want %s", got, tc.want)
			}
		})
	}
}

func BenchmarkPrepareDocumentHashSHA256(b *testing.B) {
	digest := make([]byte, 32)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := tcrsa.PrepareDocumentHash(64, crypto.SHA256, digest); err != nil {
			b.Fatal(err)
		}
	}
}
