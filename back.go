package main

import "fmt"

func existeCaminho(labirinto [][]int) bool {
    n := len(labirinto)

    // Se início ou fim forem bloqueados, não há caminho
    if labirinto[0][0] == 1 || labirinto[n-1][n-1] == 1 {
        return false
    }

    // Movimentos possíveis: direita, esquerda, baixo, cima
    movimentos := [][]int{
        {0, 1},  // direita
        {0, -1}, // esquerda
        {1, 0},  // baixo
        {-1, 0}, // cima
    }

    visitado := make([][]bool, n)
    for i := range visitado {
        visitado[i] = make([]bool, n)
    }

    var dfs func(x, y int) bool
    dfs = func(x, y int) bool {
        // Chegou ao destino
        if x == n-1 && y == n-1 {
            return true
        }

        visitado[x][y] = true

        // Tenta todos os movimentos possíveis
        for _, m := range movimentos {
            nx, ny := x+m[0], y+m[1]

            // Verifica limites
            if nx >= 0 && nx < n && ny >= 0 && ny < n {
                // Verifica se é caminho e não foi visitado
                if labirinto[nx][ny] == 0 && !visitado[nx][ny] {
                    if dfs(nx, ny) {
                        return true
                    }
                }
            }
        }
        return false
    }

    return dfs(0, 0)
}

func main() {
    labirinto := [][]int{
        {0, 1, 0, 0},
        {0, 0, 0, 1},
        {1, 1, 0, 0},
        {0, 0, 0, 0},
    }

    if existeCaminho(labirinto) {
        fmt.Println("Sim")
    } else {
        fmt.Println("Não")
    }
}
