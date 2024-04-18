package coinflip

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)

// Function to verify the game result
func VerifyResult(serverSeed, clientSeed, round string) string {

	hashStr := fmt.Sprintf("%s-%s-%s", serverSeed, clientSeed, round)

	// Compute the SHA-256 hash
	hash := sha256.Sum256([]byte(hashStr))
	hashInt := new(big.Int)
	hashInt.SetString(hex.EncodeToString(hash[:8]), 16)
	flip := hashInt.Mod(hashInt, big.NewInt(2)).Uint64() + 1

	// Output the result based on the computed hash
	if flip == 1 {
		return "Tails"
	} else {
		return "Heads"
	}
}
