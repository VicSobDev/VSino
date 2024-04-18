package roulette

import (
	"fmt"
	"math/big"

	"github.com/vicsobdev/VSino/games/utils"
)

// Function to verify the game result
func VerifyResult(serverSeed, clientSeed, round string) string {

	hash := utils.SaltHash(serverSeed, clientSeed+"-"+round)

	num := new(big.Int)
	num.SetString(hash[:52/4], 16)
	result := num.Mod(num, big.NewInt(15)).Int64()
	switch {
	case result == 0:
		return "gold"
	case result <= 7:
		return "bronze"
	default:
		fmt.Println(result)
		return "silver"
	}
}
