package main

import (
	"bufio"
	"fmt"
	"github.com/liuchibing/mathgo"
	"os"
)

const VERSION string = "0.1.2"

func main() {
	fmt.Println("Interactive-MathGo version", VERSION)
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
