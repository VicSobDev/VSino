package roulette

import (
	"fmt"
	"math/big"

	"github.com/google/uuid"
	"github.com/vicsobdev/VSino/games/utils"
)

// Game represents a roulette game instance with server, user seeds, and a unique round identifier.
type Game struct {
	serverSeed string
	userSeed   string
	round      string
}

// NewGame initializes a new Game instance with a provided round identifier.
func NewGame(round string) *Game {
	return &Game{round: round}
}

// GenServerSeed generates a unique server seed, hashes it, and returns the hash.
func (g *Game) GenServerSeed() string {
	g.serverSeed = uuid.NewString()
	return utils.Hash256(g.serverSeed)
}

// SetUserSeed sets the user's seed for the game.
func (g *Game) SetUserSeed(seed string) {
	g.userSeed = seed
}

// GenResult calculates the game's result based on the seeds and round, then returns the result and the original server seed.
func (g *Game) GenResult() (string, string) {
	hash := utils.SaltHash(g.serverSeed, g.userSeed+"-"+g.round)
	result := gameResultFromHash(hash)
	fmt.Println("Generated game result:", result)
	return result, g.serverSeed
}

// gameResultFromHash computes the roulette result from a hash value.
func gameResultFromHash(hash string) string {
	// Create a new big integer from the first 13 characters of the hash, interpreted in base 16.
	num := new(big.Int)
	num.SetString(hash[:13], 16)

	// Compute the modulo of 15 to get a result between 0 and 14.
	modResult := num.Mod(num, big.NewInt(15)).Int64()

	// Map the result to game outcomes.
	switch {
	case modResult == 0:
		return "gold"
	case modResult <= 7:
		return "bronze"
	default:
		return "silver"
	}
}
