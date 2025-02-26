package tcrsa

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"testing"
	"time"
)

const utilsTestBitlen = 256

// Miller-Rabin primality test rounds
const utilsTestC = 25

// Tests that two consecutive outputs from random dev are different.
// TODO: Test how much different are the numbers generated
func TestRandomDev_different(t *testing.T) {
	rand1, err := randInt(utilsTestBitlen)
	if err != nil {
		t.Errorf("first random number generation failed: %v", err)
	}
	rand2, err := randInt(utilsTestBitlen)
	if err != nil {
		t.Errorf("second random number generation failed: %v", err)
	}
	if rand1.Cmp(rand2) == 0 {
		t.Errorf("both random numbers are equal!")
	}
}

// Tests that the bit size of the output of a random dev function is the desired.
func TestRandomDev_bitSize(t *testing.T) {
	rand1, err := randInt(utilsTestBitlen)
	if err != nil {
		t.Errorf("first random number generation failed: %v", err)
	}
	if rand1.BitLen() > utilsTestBitlen {
		t.Errorf("random number bit length should have been at most %d, but it was %d", rand1.BitLen(), utilsTestBitlen)
	}
}

func TestGenerateSafePrimes(t *testing.T) {

	pExpected := new(big.Int)

	p, pr, err := generateSafePrimes(utilsTestBitlen, rand.Reader)
	if err != nil {
		t.Errorf("safe prime generation failed: %v", err)
	}
	if !p.ProbablyPrime(utilsTestC) {
		t.Errorf("p is not prime")
	}
	if !pr.ProbablyPrime(utilsTestC) {
		t.Errorf("pr is not prime")
	}
	pExpected.Mul(pr, big.NewInt(2)).Add(pExpected, big.NewInt(1))
	if p.Cmp(pExpected) != 0 {
		t.Errorf("p is not 2*pr + 1")
	}
}

func TestGenerateSafePrimes_keyGeneration(t *testing.T) {

	m := new(big.Int)
	d := new(big.Int)
	r := new(big.Int)

	_, pr, err := generateSafePrimes(utilsTestBitlen, rand.Reader)
	if err != nil {
		t.Errorf("safe prime generation failed: %v", err)
	}

	_, qr, err := generateSafePrimes(utilsTestBitlen, rand.Reader)
	if err != nil {
		t.Errorf("safe prime generation failed: %v", err)
	}

	m.Mul(pr, qr)
	e := big.NewInt(65537)

	d.ModInverse(e, m)
	r.Mul(d, e).Mod(r, m)

	if r.Cmp(big.NewInt(1)) != 0 {
		t.Errorf("safe prime generation failed")
	}

}

func TestGenerateSafePrimes_Time(t *testing.T) {
	for i := 4; i < 11; i++ {
		keyLength := 1 << uint(i)
		start := time.Now()
		_, _, err := generateSafePrimes(keyLength, rand.Reader)
		if err != nil {
			t.Errorf("error generating safe primes: %d", err)
		}
		fmt.Printf("- %d byte safe prime pair obtained in %f seconds\n", keyLength, time.Since(start).Seconds())
	}
}
