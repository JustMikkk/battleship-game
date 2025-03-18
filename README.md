# Battleship Game

This is a Go implementation of the classic Battleship game. The game generates a random board with ships placed on it, and the player guesses the coordinates to sink all the ships.

---

## Features

- Randomized ship placement with no overlapping or touching ships.
- Interactive gameplay with hit/miss feedback.
- Tracks sunk ships and displays remaining ships to find.
- Fully written in Go with clean and modular code.

---

## Project Structure

```
battleship-game
├── main.go        # Entry point of the application
├── main_test.go   # Unit tests for the game logic
├── README.md      # Documentation for the project
```

---

## How to Run the Game

1. **Install Go**: Ensure you have Go installed on your machine. You can download it from [golang.org](https://golang.org/).
2. **Clone the Repository**: Clone this repository or download the project files.
3. **Navigate to the Project Directory**: Open a terminal and navigate to the project folder.
4. **Run the Game**: Use the following command to start the game:
   ```bash
   go run main.go
   ```

---

## How to Play

1. **Objective**: Sink all the ships on the board by guessing their coordinates.
2. **Input Format**: Enter your guesses in the format `row,column` (e.g., `2,3`).
3. **Game Feedback**:
   - **Hit**: If your guess hits a ship, it will be marked with `#`.
   - **Miss**: If your guess misses, it will be marked with `O`.
4. **Track Progress**: The game will display the remaining ships to find after each guess.
5. **Win Condition**: The game ends when all ships are sunk.

---

## Example Gameplay

```
Welcome to Battleship!
Try to guess the ship placements on the board.
Enter your guesses in the format 'row,column' (e.g., '2,3').

Your guesses:
    0   1   2   3   4   5   6   7   8   9  
   -----------------------------------------
 0 | . | . | . | . | . | . | . | . | . | . |
   -----------------------------------------
 1 | . | . | . | . | . | . | . | . | . | . |
   -----------------------------------------
 2 | . | . | . | . | . | . | . | . | . | . |
   -----------------------------------------
 3 | . | . | . | . | . | . | . | . | . | . |
   -----------------------------------------
 4 | . | . | . | . | . | . | . | . | . | . |
   -----------------------------------------
 5 | . | . | . | . | . | . | . | . | . | . |
   -----------------------------------------
 6 | . | . | . | . | . | . | . | . | . | . |
   -----------------------------------------
 7 | . | . | . | . | . | . | . | . | . | . |
   -----------------------------------------
 8 | . | . | . | . | . | . | . | . | . | . |
   -----------------------------------------
 9 | . | . | . | . | . | . | . | . | . | . |

Enter your guess (or 'exit' to quit): 2,3
Miss!

Enter your guess (or 'exit' to quit): 4,5
Hit!

Ships left to find:
1x ####
2x ###
3x ##
4x #
```

---

## Testing the Game

To run the unit tests for the game, use the following command:
```bash
go test -v
```

---

## Future Improvements

- Add support for multiplayer mode.
- Implement a graphical user interface (GUI).
- Add difficulty levels with varying board sizes and ship counts.

---

## License

This project is open-source and available under the MIT License.

---

Enjoy playing Battleship!