package main

import (
    "testing"
    "math/rand"
)

func TestGenerateBoard(t *testing.T) {
    board, ships := generateBoard()
    
    // Test board size
    if len(board) != boardSize {
        t.Errorf("Expected board height %d, got %d", boardSize, len(board))
    }
    for _, row := range board {
        if len(row) != boardSize {
            t.Errorf("Expected board width %d, got %d", boardSize, len(row))
        }
    }
    
    // Test total ship cells
    var shipCells int
    for _, row := range board {
        for _, cell := range row {
            shipCells += cell
        }
    }
    expectedCells := 20 // Sum of shipLengths
    if shipCells != expectedCells {
        t.Errorf("Expected %d ship cells, got %d", expectedCells, shipCells)
    }
    
    // Test number of ships
    if len(ships) != len(shipLengths) {
        t.Errorf("Expected %d ships, got %d", len(shipLengths), len(ships))
    }
}

func TestPlaceShip(t *testing.T) {
    // Test valid placement
    board := make([][]int, boardSize)
    for i := range board {
        board[i] = make([]int, boardSize)
    }
    
    // Place test ship with fixed seed
    rand.Seed(1)
    ship := placeShip(board, 3)
    
    // Verify ship positions
    if len(ship.positions) != 3 {
        t.Errorf("Expected ship length 3, got %d", len(ship.positions))
    }
    
    // Create a map of ship positions for quick lookup
    shipPositions := make(map[[2]int]bool)
    for _, pos := range ship.positions {
        shipPositions[[2]int{pos[0], pos[1]}] = true
    }
    
    // Verify board markings and no adjacent ships
    for _, pos := range ship.positions {
        row, col := pos[0], pos[1]
        if board[row][col] != 1 {
            t.Errorf("Expected 1 at (%d, %d), got %d", row, col, board[row][col])
        }
        
        // Check surrounding cells (excluding ship's own positions)
        for dr := -1; dr <= 1; dr++ {
            for dc := -1; dc <= 1; dc++ {
                if dr == 0 && dc == 0 {
                    continue // Skip current cell
                }
                nr, nc := row+dr, col+dc
                if nr >= 0 && nr < boardSize && nc >= 0 && nc < boardSize {
                    // Skip checking other parts of the same ship
                    if shipPositions[[2]int{nr, nc}] {
                        continue
                    }
                    if board[nr][nc] == 1 {
                        t.Errorf("Found adjacent ship at (%d, %d)", nr, nc)
                    }
                }
            }
        }
    }
}

func TestParseGuess(t *testing.T) {
    tests := []struct {
        input    string
        expectedRow int
        expectedCol int
        shouldErr  bool
    }{
        {"0,0", 0, 0, false},
        {"5,5", 5, 5, false},
        {"9,9", 9, 9, false},
        {"10,5", 0, 0, true},    // Invalid row
        {"5,10", 0, 0, true},    // Invalid column
        {"-1,3", 0, 0, true},    // Negative row
        {"2,-3", 0, 0, true},    // Negative column
        {"abc,3", 0, 0, true},   // Invalid format
        {"5", 0, 0, true},       // Missing comma
        {"3,4,5", 0, 0, true},   // Too many parts
    }

    for _, tt := range tests {
        row, col, err := parseGuess(tt.input)
        if tt.shouldErr {
            if err == nil {
                t.Errorf("Expected error for input %q, got none", tt.input)
            }
        } else {
            if err != nil {
                t.Errorf("Unexpected error for input %q: %v", tt.input, err)
            }
            if row != tt.expectedRow || col != tt.expectedCol {
                t.Errorf("For input %q, expected (%d,%d), got (%d,%d)", 
                    tt.input, tt.expectedRow, tt.expectedCol, row, col)
            }
        }
    }
}

func TestHitDetection(t *testing.T) {
    // Create test board with known ship positions
    board := make([][]int, boardSize)
    for i := range board {
        board[i] = make([]int, boardSize)
    }
    ship := Ship{
        positions: [][2]int{{2, 2}, {2, 3}, {2, 4}},
        length:    3,
    }
    for _, pos := range ship.positions {
        board[pos[0]][pos[1]] = 1
    }

    // Test hit
    row, col := 2, 3
    if board[row][col] != 1 {
        t.Errorf("Expected hit at (%d, %d)", row, col)
    }

    // Test miss
    row, col = 5, 5
    if board[row][col] != 0 {
        t.Errorf("Expected miss at (%d, %d)", row, col)
    }
}

func TestSunkShip(t *testing.T) {
    // Create test ships
    ships := []Ship{
        {
            positions: [][2]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
            length:    4,
        },
        {
            positions: [][2]int{{5, 5}},
            length:    1,
        },
    }

    // Create guesses that sink the large ship
    guesses := make([][]string, boardSize)
    for i := range guesses {
        guesses[i] = make([]string, boardSize)
        for j := range guesses[i] {
            guesses[i][j] = "."
        }
    }
    
    // Mark all positions of the first ship as hit
    for _, pos := range ships[0].positions {
        guesses[pos[0]][pos[1]] = "#"
    }

    // Verify ship sinking
    remaining := make(map[int]int)
    for _, ship := range ships {
        remaining[ship.length]++
    }

    for _, ship := range ships {
        sunk := true
        for _, pos := range ship.positions {
            if guesses[pos[0]][pos[1]] != "#" {
                sunk = false
                break
            }
        }
        if sunk {
            remaining[ship.length]--
        }
    }

    if remaining[4] != 0 {
        t.Errorf("Expected 4-length ship to be sunk, got %d remaining", remaining[4])
    }
    if remaining[1] != 1 {
        t.Errorf("Expected 1 1-length ship remaining, got %d", remaining[1])
    }
}

func TestEdgePlacements(t *testing.T) {
    board := make([][]int, boardSize)
    for i := range board {
        board[i] = make([]int, boardSize)
    }

    // Test vertical ship at bottom edge
    if !canPlaceShip(board, boardSize-3, 5, 3, 1) {
        t.Error("Should be able to place vertical ship at bottom edge")
    }

    // Test horizontal ship at right edge
    if !canPlaceShip(board, 5, boardSize-3, 3, 0) {
        t.Error("Should be able to place horizontal ship at right edge")
    }

    // Test invalid placement beyond edges
    if canPlaceShip(board, boardSize-2, 5, 3, 1) {
        t.Error("Should not be able to place vertical ship beyond bottom edge")
    }
    
    if canPlaceShip(board, 5, boardSize-2, 3, 0) {
        t.Error("Should not be able to place horizontal ship beyond right edge")
    }
}