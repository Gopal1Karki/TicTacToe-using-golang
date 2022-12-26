package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

//front page setup

func frontPage() {

RES:
	clearScreen()
	fmt.Println("######################")
	fmt.Println("#    TIC TAC TOE     #")
	fmt.Println("######################")
	fmt.Println("Press c to continue the game.")
	var n string
	fmt.Scanln(&n)
	switch {
	case n == "c":
		clearScreen()
		playTicTacToe()
	default:
		fmt.Println("Invalid Input")
		goto RES
	}
}

// Code for clearing the terminal
func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// design of TicTacToe
func playTicTacToe() {
	len := 3
	var design = make([]string, len)
	for i := 0; i < 3; i++ {
		design[i] = " "
		fmt.Printf("|")
		fmt.Printf(design[i])
		fmt.Printf("|")
	}
	fmt.Println("")
	var design1 = make([]string, len)
	for i := 0; i < 3; i++ {
		design1[i] = " "
		fmt.Printf("|")
		fmt.Printf(design[i])
		fmt.Printf("|")
	}
	fmt.Println("")
	var design2 = make([]string, len)
	for i := 0; i < 3; i++ {
		design2[i] = " "
		fmt.Printf("|")
		fmt.Printf(design[i])
		fmt.Printf("|")
	}
	fmt.Println("")
	count := 0
	fmt.Println("Enter the value from 1 - 9 to fill the value. ")
LOOP:

	//var player1 = 0
	var inp int
	count = count + 1
	if count > 9 {
		fmt.Println("The game is draw.")
		gameOver()
	}
	if count == 1 || count == 3 || count == 5 || count == 7 || count == 9 {
		fmt.Println("Player One Turn")
		fmt.Println("Enter the position Player One:")
	} else if count == 2 || count == 4 || count == 6 || count == 8 {
		fmt.Println("Player Two Turn")
		fmt.Println("Enter the position Player Two:")
	}
GOP:
	fmt.Scanln(&inp)
	if inp < 1 || inp > 9 {
		fmt.Println("Enter the valid position from 1-9 ")
		fmt.Println("Enter the position again.")
		goto GOP
	}
	if inp <= 3 {
		if design[inp-1] == "X" || design[inp-1] == "O" {
			fmt.Println("The value you have inserted is already taken position")
			fmt.Println("Enter the position again.")
			goto GOP
		}
	}
	if inp > 3 && inp <= 6 {
		if (design1[inp-4] == "X") || design1[inp-4] == "O" {
			fmt.Println("The value you have inserted is already taken position")
			fmt.Println("Enter the position again.")
			goto GOP
		}
	}
	if inp > 6 && inp <= 9 {
		if design2[inp-7] == "X" || design2[inp-7] == "O" {
			fmt.Println("The value you have inserted is already taken position")
			fmt.Println("Enter the position again.")
			goto GOP
		}
	}

	a := inp
	b := count
	switch count {
	case 1, 3, 5, 7, 9:
		fmt.Printf("Player 1 entered postion %d \n", inp)
		if b == 1 || b == 3 || b == 5 || b == 7 || b == 9 {
			if a == 1 || a == 2 || a == 3 {
				design[a-1] = "O"
			} else if a == 4 || a == 5 || a == 6 {
				design1[a-4] = "O"
			} else if a == 7 || a == 8 || a == 9 {
				design2[a-7] = "O"
			}
		}
		clearScreen()
		fmt.Println(design)
		fmt.Println(design1)
		fmt.Println(design2)
		time.Sleep(2 * time.Second)
		if design[0] == design[1] && design[1] == design[2] && design[0] == "O" && design[1] == "/0" {
			fmt.Println("Player One wins")
			gameOver()
		}
		if design1[0] == design1[1] && design1[0] == design1[2] && design1[0] == "O" {
			fmt.Println("Player One wins")
			gameOver()
		}
		if design2[0] == design2[1] && design2[0] == design2[2] && design2[0] == "O" {
			fmt.Println("Player One wins")
			gameOver()
		}
		if design[0] == design1[0] && design[0] == design2[0] && design[0] == "O" {
			fmt.Println("Player One wins")
			gameOver()
		}
		if design[1] == design1[1] && design[1] == design2[1] && design[1] == "O" {
			fmt.Println("Player One wins")
			gameOver()
		}
		if design[2] == design1[2] && design[2] == design2[2] && design[2] == "O" {
			fmt.Println("Player One wins")
			gameOver()
		}
		if design[0] == design1[1] && design[0] == design2[2] && design[0] == "O" {
			fmt.Println("Player One wins")
			gameOver()
		}
		if design[2] == design1[1] && design[2] == design2[0] && design[2] == "O" {
			fmt.Println("Player One wins")
			gameOver()
		}
		goto LOOP
	case 2, 4, 6, 8:
		fmt.Printf("Player 2 entered postion %d \n", inp)
		if b == 2 || b == 4 || b == 6 || b == 8 {
			if a == 1 || a == 2 || a == 3 {
				design[a-1] = "X"
			} else if a == 4 || a == 5 || a == 6 {
				design1[a-4] = "X"
			} else if a == 7 || a == 8 || a == 9 {
				design2[a-7] = "X"
			}
		}
		clearScreen()
		fmt.Println(design)
		fmt.Println(design1)
		fmt.Println(design2)
		time.Sleep(2 * time.Second)
		if design[0] == design[1] && design[0] == design[2] && design[0] == "X" {
			fmt.Println("Player Two wins")
			gameOver()
		}
		if design1[0] == design1[1] && design1[0] == design1[2] && design1[0] == "X" {
			fmt.Println("Player Two wins")
			gameOver()
		}
		if design2[0] == design2[1] && design2[0] == design2[2] && design[0] == "X" {
			fmt.Println("Player Two wins")
			gameOver()
		}
		if design[0] == design1[0] && design[0] == design2[0] && design[0] == "X" {
			fmt.Println("Player Two wins")
			gameOver()
		}
		if design[1] == design1[1] && design[1] == design2[1] && design[1] == "X" {
			fmt.Println("Player Two wins")
			gameOver()
		}
		if design[2] == design1[2] && design[2] == design2[2] && design[2] == "X" {
			fmt.Println("Player Two wins")
			gameOver()
		}
		if design[0] == design1[1] && design[0] == design2[2] && design[0] == "X" {
			fmt.Println("Player Twp wins")
			gameOver()
		}
		if design[2] == design1[1] && design[2] == design2[0] && design2[0] == "X" {
			fmt.Println("Player Two wins")
			gameOver()
		}
		goto LOOP
	default:
		fmt.Println("Invalid Input. Enter the valid input of rage 1-9")
		goto LOOP
	}
}

// gameOver
func gameOver() {
	fmt.Println("--------------------------")
	fmt.Println("|    G A M E  O V E R    |")
	fmt.Println("--------------------------")
	time.Sleep(5 * time.Second)
REC:
	var xy string
	fmt.Println("Press c to continue or e to exit the game....")
	fmt.Scanln(&xy)
	if xy == "c" {
		playTicTacToe()
	} else if xy == "e" {
		os.Exit(5)
	} else {
		fmt.Println("Enter the valid Input !!")
		goto REC
	}
}

// main function of the game
func main() {
	frontPage()
}
