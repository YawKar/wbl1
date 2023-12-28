package _15

import "testing"

func TestProofMemoryLeakage(t *testing.T) {
	if err := ProofStringLeakageFunc(); err != nil {
		t.Fatalf("proof failed: %s", err)
	}
}

func TestProofNoMemoryLeakage(t *testing.T) {
	if err := ProofWithoutLeakageCorrectness(); err != nil {
		t.Fatalf("proof failed: %s", err)
	}
}
