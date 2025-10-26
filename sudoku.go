package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const N = 9

type Grid [N][N]int

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("=== Sudoku en Go ===")
	fmt.Println("Choisis une difficulté : easy / medium / hard")
	fmt.Print("> ")

	reader := bufio.NewReader(os.Stdin)
	diff, _ := reader.ReadString('\n')
	diff = strings.TrimSpace(strings.ToLower(diff))

	cellsToRemove := 45 // par défaut moyen
	switch diff {
	case "easy":
		cellsToRemove = 35
	case "hard":
		cellsToRemove = 55
	default:
		fmt.Println("Difficulté par défaut : medium")
	}

	// Génère une grille complète et le puzzle
	solution := Grid{}
	generateFull(&solution)

	puzzle := solution
	removeCells(&puzzle, cellsToRemove)

	givens := [N][N]bool{}
	for r := 0; r < N; r++ {
		for c := 0; c < N; c++ {
			if puzzle[r][c] != 0 {
				givens[r][c] = true
			}
		}
	}

	for {
		printGrid(puzzle)
		fmt.Println("Commandes : `ligne colonne valeur` (ex: 1 3 9), `hint`, `solve`, `quit`.")
		fmt.Print("> ")
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		switch strings.ToLower(line) {
		case "quit", "exit", "q":
			fmt.Println("À bientôt 👋")
			return
		case "solve":
			fmt.Println("Solution :")
			printGrid(solution)
			return
		case "hint":
			if giveHint(&puzzle, &solution, &givens) {
				fmt.Println("Indice ajouté !")
				if isComplete(puzzle) {
					fmt.Println("Bravo 🎉 tu as terminé le Sudoku !")
					printGrid(puzzle)
					return
				}
			} else {
				fmt.Println("Aucun indice disponible.")
			}
			continue
		default:
			parts := strings.Fields(line)
			if len(parts) != 3 {
				fmt.Println("Entrée invalide.")
				continue
			}
			r, _ := strconv.Atoi(parts[0])
			c, _ := strconv.Atoi(parts[1])
			v, _ := strconv.Atoi(parts[2])

			if r < 1 || r > 9 || c < 1 || c > 9 || v < 0 || v > 9 {
				fmt.Println("Lignes/colonnes: 1–9. Valeur: 0–9 (0 pour effacer).")
				continue
			}
			r--
			c--

			if givens[r][c] {
				fmt.Println("Impossible de modifier une case donnée.")
				continue
			}

			if v == 0 {
				puzzle[r][c] = 0
				continue
			}

			if isValidMove(&puzzle, r, c, v) {
				puzzle[r][c] = v
				if isComplete(puzzle) {
					fmt.Println("Bravo 🎉 Sudoku complété !")
					printGrid(puzzle)
					return
				}
			} else {
				fmt.Println("Mouvement invalide (règle du Sudoku).")
			}
		}
	}
}

// --- Génération et logique ---

func printGrid(g Grid) {
	fmt.Println()
	fmt.Println("   1 2 3   4 5 6   7 8 9")
	for r := 0; r < N; r++ {
		if r%3 == 0 {
			fmt.Println("  +-------+-------+-------+")
		}
		fmt.Printf("%d |", r+1)
		for c := 0; c < N; c++ {
			if g[r][c] == 0 {
				fmt.Print(" .")
			} else {
				fmt.Printf(" %d", g[r][c])
			}
			if (c+1)%3 == 0 {
				fmt.Print(" |")
			}
		}
		fmt.Println()
	}
	fmt.Println("  +-------+-------+-------+")
	fmt.Println()
}

func generateFull(g *Grid) bool {
	for r := 0; r < N; r++ {
		for c := 0; c < N; c++ {
			g[r][c] = 0
		}
	}
	return fillGrid(g)
}

func fillGrid(g *Grid) bool {
	r, c, found := findEmpty(g)
	if !found {
		return true
	}
	nums := rand.Perm(9)
	for _, n := range nums {
		val := n + 1
		if isValidMove(g, r, c, val) {
			g[r][c] = val
			if fillGrid(g) {
				return true
			}
			g[r][c] = 0
		}
	}
	return false
}

func findEmpty(g *Grid) (int, int, bool) {
	for r := 0; r < N; r++ {
		for c := 0; c < N; c++ {
			if g[r][c] == 0 {
				return r, c, true
			}
		}
	}
	return -1, -1, false
}

func isValidMove(g *Grid, r, c, val int) bool {
	for i := 0; i < N; i++ {
		if g[r][i] == val || g[i][c] == val {
			return false
		}
	}
	br, bc := (r/3)*3, (c/3)*3
	for i := br; i < br+3; i++ {
		for j := bc; j < bc+3; j++ {
			if g[i][j] == val {
				return false
			}
		}
	}
	return true
}

func removeCells(g *Grid, k int) {
	positions := rand.Perm(81)
	removed := 0
	for _, pos := range positions {
		if removed >= k {
			break
		}
		r := pos / 9
		c := pos % 9
		if g[r][c] == 0 {
			continue
		}
		backup := g[r][c]
		g[r][c] = 0

		var copyGrid Grid = *g
		count := 0
		countSolutions(&copyGrid, &count, 2)
		if count != 1 {
			g[r][c] = backup
		} else {
			removed++
		}
	}
}

func countSolutions(g *Grid, count *int, limit int) {
	if *count >= limit {
		return
	}
	r, c, found := findEmpty(g)
	if !found {
		*count++
		return
	}
	for val := 1; val <= 9; val++ {
		if isValidMove(g, r, c, val) {
			g[r][c] = val
			countSolutions(g, count, limit)
			g[r][c] = 0
			if *count >= limit {
				return
			}
		}
	}
}

func giveHint(puzzle *Grid, solution *Grid, givens *[N][N]bool) bool {
	for r := 0; r < N; r++ {
		for c := 0; c < N; c++ {
			if puzzle[r][c] == 0 && !givens[r][c] {
				puzzle[r][c] = solution[r][c]
				return true
			}
		}
	}
	return false
}

func isComplete(g Grid) bool {
	for r := 0; r < N; r++ {
		for c := 0; c < N; c++ {
			if g[r][c] == 0 {
				return false
			}
		}
	}
	return true
}
