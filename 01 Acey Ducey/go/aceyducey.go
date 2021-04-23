package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//N:=100
	Q := 100
	var A, B, C, M int
	var S string
	rand.Seed(time.Now().UnixNano())
	fmt.Println("\t\t\tACEY DUCEY CARD GAME")
	fmt.Println("\t\tCREATIVE COMPUTING  MORRISTOWN, NEW JERSEY")
	fmt.Println()
	fmt.Println()
	fmt.Println("ACEY-DUCEY IS PLAYED IN THE FOLLOWING MANNER ")
	fmt.Println("THE DEALER (COMPUTER) DEALS TWO CARDS FACE UP")
	fmt.Println("YOU HAVE AN OPTION TO BET OR NOT BET DEPENDING")
	fmt.Println("ON WHETHER OR NOT YOU FEEL THE CARD WILL HAVE")
	fmt.Println("A VALUE BETWEEN THE FIRST TWO.")
	fmt.Println("IF YOU DO NOT WANT TO BET, INPUT A 0")
	fmt.Println()
	fmt.Println("YOU NOW HAVE", Q, "DOLLARS.")
	fmt.Println()
	for {
		fmt.Println("HERE ARE YOUR NEXT TWO CARDS: ")
		//[0,13) +2
		//0,1,2,3,4,5,6,7,8,9,10,11,12
		//2,3,4,5,6,7,8,9,10,J,Q,K,A
		A = rand.Intn(13) + 2
		B = rand.Intn(13) + 2
		for A >= B {
			A = rand.Intn(13) + 2
			B = rand.Intn(13) + 2
		}

		if A < 11 {
			fmt.Println(A)
		} else if A == 11 {
			fmt.Println("JACK")
		} else if A == 12 {
			fmt.Println("QUEEN")
		} else if A == 13 {
			fmt.Println("KING")
		} else if A == 14 {
			fmt.Println("ACE")
		}

		if B < 11 {
			fmt.Println(B)
		} else if B == 11 {
			fmt.Println("JACK")
		} else if B == 12 {
			fmt.Println("QUEEN")
		} else if B == 13 {
			fmt.Println("KING")
		} else if B == 14 {
			fmt.Println("ACE")
		}

		for {
			fmt.Println()
			fmt.Println("WHAT IS YOUR BET")
			fmt.Scanf("%d", &M)
			if M > Q {
				fmt.Println("SORRY, MY FRIEND, BUT YOU BET TOO MUCH.")
				fmt.Println("YOU HAVE ONLY", Q, "DOLLARS TO BET.")
				continue
			}
			break
		}

		if M == 0 {
			fmt.Println("CHICKEN!!")
			fmt.Println()
			continue
		}

		C = rand.Intn(13) + 2
		fmt.Println()
		if C < 11 {
			fmt.Println(C)
		} else if C == 11 {
			fmt.Println("JACK")
		} else if C == 12 {
			fmt.Println("QUEEN")
		} else if C == 13 {
			fmt.Println("KING")
		} else if C == 14 {
			fmt.Println("ACE")
		}

		fmt.Println()

		if C > A && C < B {
			fmt.Println("YOU WIN!!!")
			Q = Q + M
		} else {
			fmt.Println("SORRY, YOU LOSE")
			Q = Q - M
		}

		if Q <= 0 {
			fmt.Println()
			fmt.Println("SORRY FRIEND, BUT YOU BLEW YOUR WAD.")
			fmt.Println()
			fmt.Println("TRY AGAIN (YES OR NO)")
			fmt.Scanf("%s", &S)
			if S == "YES" {
				Q = 100
				fmt.Println()
				fmt.Println("YOU NOW HAVE", Q, "DOLLARS.")
				continue
			} else {
				fmt.Println()
				fmt.Println("O.K., HOPE YOU HAD FUN!")
				fmt.Println()
				break
			}
		}

		fmt.Println()
		fmt.Println("YOU NOW HAVE", Q, "DOLLARS.")
	}
}
