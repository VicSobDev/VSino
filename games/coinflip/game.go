package coinflip

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/google/uuid"
	"github.com/vicsobdev/VSino/games/utils"
)

// Game represents a coin flip game instance with server, user seeds, and a unique round identifier.
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

// Flip performs a coin flip and returns the result along with the server seed.
func (g *Game) Flip() (result string, serverSeed string) {
	// Create a hash from the concatenated server seed, user seed, and round.
	hash := sha256.Sum256([]byte(fmt.Sprintf("%s-%s-%s", g.serverSeed, g.userSeed, g.round)))

	// Convert the first 8 bytes of the hash to an integer.
	hashInt := new(big.Int)
	hashInt.SetString(hex.EncodeToString(hash[:8]), 16)

	// Determine the coin flip result by modulo operation on 2, results are either 0 (Tails) or 1 (Heads).
	flipResult := hashInt.Mod(hashInt, big.NewInt(2)).Uint64()

	if flipResult == 0 {
		result = "Tails"
	} else {
		result = "Heads"
	}

	return result, g.serverSeed
}
