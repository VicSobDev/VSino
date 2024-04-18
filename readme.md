# VSino Games Proof of Concept

Welcome to the VSino Games repository. This repository contains the simplified implementation of the "probably fair" system as a proof of concept (PoC) for a gambling game system. It includes basic game mechanics for roulette and coin flip games.

## Project Overview

This project was created months ago as a proof of concept for another project. The current implementation serves to demonstrate the fundamental mechanics behind a fairness system in online gambling games.

### Disclaimer

This code is a **simplified version** of a "probably fair" system. It has known weaknesses and should **not** be used in production environments. The current approach has several limitations:

- The server seed generation is simplistic and only hashed once, which might be vulnerable to predictability in a real-world scenario.
- A more robust implementation would involve hashing the server seed multiple times and potentially using a more secure algorithm to generate the server seed.

### Included Games

- **Coin Flip**: A simple game where a coin is flipped to return either "Heads" or "Tails".
- **Roulette**: A basic implementation of the roulette game where the result is determined by a hashed seed.

## Getting Started

To run these games on your local machine, follow these steps:

1. Clone this repository:
   ```bash
   git clone https://github.com/VicSobDev/VSino
   ```
2. Navigate into the project directory:
   ```bash
   cd VSino
   ```
3. To run the games, execute:
   ```bash
   go run main.go
   ```