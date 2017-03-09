package main

import (
	"bufio"
	"fmt"
	"github.com/liuchibing/mathgo"
	"os"
)

func main() {
	fmt.Println("Interactive-MathGo version 1.0")
	inter := mathgo.CreateInterpreter()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-->")
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		if input == "exit\n" {
			return
		}

		res, needNext := inter.Run(input)
		if needNext {
			continue
		}

		fmt.Println("=>", res)
		fmt.Print(input)
		fmt.Print("-->")
	}
}
