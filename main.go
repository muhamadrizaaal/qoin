package main

import (
    "fmt"
    "math/rand"
    "time"
)

// struktur untuk pemain
type Player struct {
    id    int
    dice  []int
    points int
}

// fungsi untuk melempar dadu
func rollDice(n int) []int {
    dice := make([]int, n)
    for i := 0; i < n; i++ {
        dice[i] = rand.Intn(6) + 1
    }
    return dice
}

// fungsi untuk evaluasi hasil lemparan dadu
func evaluateDice(players []*Player) {
    for i := range players {
        newDice := []int{}
        for _, die := range players[i].dice {
            switch die {
            case 6:
                players[i].points++
            case 1:
                nextPlayer := (i + 1) % len(players)
                players[nextPlayer].dice = append(players[nextPlayer].dice, 1)
            default:
                newDice = append(newDice, die)
            }
        }
        players[i].dice = newDice
    }
}

// fungsi utama untuk menjalankan permainan
func main() {
    var N, M int
    fmt.Print("Masukkan jumlah pemain: ")
    fmt.Scan(&N)
    fmt.Print("Masukkan jumlah dadu per pemain: ")
    fmt.Scan(&M)

    // inisialisasi pemain
    players := make([]*Player, N)
    for i := 0; i < N; i++ {
        players[i] = &Player{
            id:   i + 1,
            dice: rollDice(M),
        }
    }

    round := 1
    rand.Seed(time.Now().UnixNano())

    for {
        fmt.Printf("==================\n")
        fmt.Printf("Giliran %d lempar dadu:\n", round)
        activePlayers := 0

        // Lempar dadu dan tampilkan hasilnya
        for _, player := range players {
            if len(player.dice) > 0 {
                activePlayers++
                fmt.Printf("Pemain #%d (%d): ", player.id, player.points)
                player.dice = rollDice(len(player.dice))
                for _, die := range player.dice {
                    fmt.Printf("%d,", die)
                }
                fmt.Println()
            }
        }

        if activePlayers <= 1 {
            break
        }

        // evaluasi hasil lemparan dadu
        evaluateDice(players)
        
        // tampilkan hasil evaluasi
        fmt.Println("Setelah evaluasi:")
        for _, player := range players {
            fmt.Printf("Pemain #%d (%d): ", player.id, player.points)
            for _, die := range player.dice {
                fmt.Printf("%d,", die)
            }
            fmt.Println()
        }
        
        round++
    }

    fmt.Printf("==================\n")
    fmt.Println("Game berakhir karena hanya satu pemain yang tersisa.")
    
    // menentukan pemenang
    winner := players[0]
    for _, player := range players {
        if player.points > winner.points {
            winner = player
        }
    }

    fmt.Printf("Game dimenangkan oleh pemain #%d dengan %d poin.\n", winner.id, winner.points)
}
