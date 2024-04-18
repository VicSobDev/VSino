package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/vicsobdev/VSino/games/coinflip"
	"github.com/vicsobdev/VSino/games/roulette"
)

func main() {
	// Print a welcome message to the user.
	fmt.Println("Welcome to VSino")

	// Display the games menu to the user.
	gamesMenu()

	// Store the user's choice of game mode.
	var gameMode string
	fmt.Scanln(&gameMode)
	CallClear()

	// Handle game mode selection using a switch statement.
	switch gameMode {
	case "1":
		// Clear the console and start the Coin Flip game.
		CallClear()
		playCoinFlip()
	case "2":
		// Clear the console and start the Roulette game.
		CallClear()
		playRoulette()
	case "0":
		// Handle verification of game results.
		handleVerification()
	default:
		// Inform the user of an invalid selection.
		fmt.Println("Invalid option selected")
	}
}

// Handle the verification of game results.
func handleVerification() {
	// Variables to store seeds and round information.
	var serverSeed, clientSeed, round string

	fmt.Print("Enter server seed: ")
	fmt.Scanln(&serverSeed)
	fmt.Print("Enter client seed: ")
	fmt.Scanln(&clientSeed)
	fmt.Print("Enter round: ")
	fmt.Scanln(&round)

	CallClear()

	// Prompt user to specify which game was played.
	var game int
	fmt.Println("What game was played?")
	fmt.Println("1. Coin Flip")
	fmt.Println("2. Roulette")

	fmt.Scanln(&game)

	// Verify and display results based on the game played.
	switch game {
	case 1:
		result := coinflip.VerifyResult(serverSeed, clientSeed, round)
		fmt.Println("Result: ", result)
	case 2:
		result := roulette.VerifyResult(serverSeed, clientSeed, round)
		fmt.Println("Result: ", result)
	}
}

// Function to play the Coin Flip game.
func playCoinFlip() {
	// Generate a new unique round ID.
	round := uuid.NewString()
	fmt.Println("Round ID: ", round)

	// Initialize the coin flip game with the round ID.
	game := coinflip.NewGame(round)
	serverHash := game.GenServerSeed()
	fmt.Println("Server Hash: ", serverHash)

	// Prompt the user to enter their seed.
	var userSeed string
	fmt.Print("Enter your seed: ")
	fmt.Scanln(&userSeed)
	game.SetUserSeed(userSeed)

	// Flip the coin and display the result.
	result, serverSeed := game.Flip()
	fmt.Println("Result: ", result)
	fmt.Println("Server Seed: ", serverSeed)
}

// Function to play Roulette.
func playRoulette() {
	// Generate a new unique round ID.
	round := uuid.NewString()
	fmt.Println("Round ID: ", round)

	// Initialize the roulette game with the round ID.
	game := roulette.NewGame(round)
	serverHash := game.GenServerSeed()
	fmt.Println("Server Hash: ", serverHash)

	// Prompt the user to enter their seed.
	var userSeed string
	fmt.Print("Enter your seed: ")
	fmt.Scanln(&userSeed)
	game.SetUserSeed(userSeed)

	// Generate the result and display it.
	result, serverSeed := game.GenResult()
	fmt.Println("Result: ", result)
	fmt.Println("Server Seed: ", serverSeed)
}
