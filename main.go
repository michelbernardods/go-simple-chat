package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/TwinProduction/go-color"
)

var user01, user02, textUser01, textUser02 string

func cadastroUser() {
	fmt.Println("Cadastro de usuário")

	fmt.Print("Nome do primeiro usuario: ")
	fmt.Scan(&user01)
	fmt.Print("Nome do segundo usuario: ")
	fmt.Scan(&user02)

	fmt.Print("\033[H\033[2J")
}

func chat() {
	fmt.Println(color.Ize(color.Red, "PARA SAIR DIGITE: exit\n"))
	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(color.Ize(color.Green, user01+": "))
		s.Scan()
		textUser01 = s.Text()
		if textUser01 == "exit" {
			fmt.Print("\033[H\033[2J")
			fmt.Println("\nSee you")
			return
		}
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("R:", color.Ize(color.Blue, textUser01))

		fmt.Print(color.Ize(color.Green, user02+": "))
		s.Scan()
		textUser02 = s.Text()
		if textUser02 == "exit" {
			fmt.Print("\033[H\033[2J")
			fmt.Println("\nSee you")
			return
		}
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("R:", color.Ize(color.Blue, textUser02))
	}
}

func main() {
	cadastroUser()
	chat()
}
