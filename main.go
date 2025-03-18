package main

import (
    "errors"
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "strings"
    "time"
)

const (
    boardSize = 10
)

var shipLengths = []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1} // Ship lengths

// Ship struct to track ship positions and length
type Ship struct {
    positions [][2]int // Each position is a [row, col] pair
    length    int
}

func main() {
    rand.Seed(time.Now().UnixNano())

    board, ships := generateBoard()
    saveBoardToFile("computer_board.txt", board)

    guesses := make([][]string, boardSize)
    for i := range guesses {
        guesses[i] = make([]string, boardSize)
        for j := range guesses[i] {
            guesses[i][j] = "."
        }
    }

    fmt.Println("Welcome to Battleship!")
    fmt.Println("Try to guess the ship placements on the board.")
    fmt.Println("Enter your guesses in the format 'row,column' (e.g., '2,3').")

    var guess string
    for {
        printGuesses(guesses)
        fmt.Print("Enter your guess (or 'exit' to quit): ")
        fmt.Scanln(&guess)

        if guess == "exit" {
            fmt.Println("Thanks for playing!")
            break
        }

        row, col, err := parseGuess(guess)
        if err != nil {
            fmt.Println("Invalid input. Please enter a valid guess.")
            continue
        }

        if guesses[row][col] != "." {
            fmt.Println("You already guessed that spot!")
            continue
        }

        if board[row][col] == 1 {
            fmt.Println("Hit!")
            guesses[row][col] = "#"
        } else {
            fmt.Println("Miss!")
            guesses[row][col] = "O"
        }

        showGuessedShips(guesses, ships)
    }
}

// Generate the board and return it along with the list of ships
func generateBoard() ([][]int, []Ship) {
    board := make([][]int, boardSize)
    for i := range board {
        board[i] = make([]int, boardSize)
    }

    var ships []Ship
    for _, length := range shipLengths {
        ship := placeShip(board, length)
        ships = append(ships, ship)
    }

    return board, ships
}

// Place a ship on the board and return its positions
func placeShip(board [][]int, length int) Ship {
    for {
        row := rand.Intn(boardSize)
        col := rand.Intn(boardSize)
        direction := rand.Intn(2) // 0 = horizontal, 1 = vertical

        if canPlaceShip(board, row, col, length, direction) {
            positions := make([][2]int, length)
            for i := 0; i < length; i++ {
                if direction == 0 {
                    board[row][col+i] = 1
                    positions[i] = [2]int{row, col + i}
                } else {
                    board[row+i][col] = 1
                    positions[i] = [2]int{row + i, col}
                }
            }
            return Ship{positions: positions, length: length}
        }
    }
}

// Check if a ship can be placed at the given position
func canPlaceShip(board [][]int, row, col, length, direction int) bool {
    for i := 0; i < length; i++ {
        r, c := row, col
        if direction == 0 {
            c += i // Horizontal
        } else {
            r += i // Vertical
        }

        // Check if the cell is out of bounds or already occupied
        if r < 0 || r >= boardSize || c < 0 || c >= boardSize {
            return false
        }
        if board[r][c] == 1 {
            return false
        }

        // Check surrounding cells to ensure no ships are touching
        for dr := -1; dr <= 1; dr++ {
            for dc := -1; dc <= 1; dc++ {
                nr, nc := r+dr, c+dc
                if nr >= 0 && nr < boardSize && nc >= 0 && nc < boardSize && board[nr][nc] == 1 {
                    return false
                }
            }
        }
    }
    return true
}

// Save the board to a file
func saveBoardToFile(filename string, board [][]int) {
    file, err := os.Create(filename)
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    for _, row := range board {
        for _, cell := range row {
            file.WriteString(fmt.Sprintf("%d ", cell))
        }
        file.WriteString("\n")
    }
}

// Parse the player's guess
func parseGuess(guess string) (int, int, error) {
    parts := strings.Split(guess, ",")
    if len(parts) != 2 {
        return 0, 0, errors.New("invalid format")
    }

    row, err := strconv.Atoi(parts[0])
    if err != nil || row < 0 || row >= boardSize {
        return 0, 0, errors.New("invalid row")
    }

    col, err := strconv.Atoi(parts[1])
    if err != nil || col < 0 || col >= boardSize {
        return 0, 0, errors.New("invalid column")
    }

    return row, col, nil
}

// Print the player's guesses
func printGuesses(guesses [][]string) {
    fmt.Println("\nYour guesses:")
    fmt.Print("    ")
    for i := 0; i < boardSize; i++ {
        fmt.Printf("%2d  ", i)
    }
    fmt.Println()

    fmt.Print("   ")
    for i := 0; i < boardSize; i++ {
        fmt.Print("----")
    }
    fmt.Println("-")

    for i, row := range guesses {
        fmt.Printf("%2d |", i)
        for _, cell := range row {
            fmt.Printf(" %s |", cell)
        }
        fmt.Println()

        fmt.Print("   ")
        for j := 0; j < boardSize; j++ {
            fmt.Print("----")
        }
        fmt.Println("-")
    }
}

// Check which ships have been sunk
func showGuessedShips(guesses [][]string, ships []Ship) {
    remainingShips := make(map[int]int)
    for _, ship := range ships {
        remainingShips[ship.length]++
    }

    for _, ship := range ships {
        sunk := true
        for _, pos := range ship.positions {
            row, col := pos[0], pos[1]
            if guesses[row][col] != "#" {
                sunk = false
                break
            }
        }
        if sunk {
            remainingShips[ship.length]--
            fmt.Printf("Ship of length %d sunk!\n", ship.length)
        }
    }

    fmt.Println("\nShips left to find:")
    for length, count := range remainingShips {
        if count > 0 {
            fmt.Printf("%dx %s\n", count, strings.Repeat("#", length))
        }
    }
}