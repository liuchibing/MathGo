package main

import (
	"bufio"
	"fmt"
	"github.com/liuchibing/mathgo"
	"os"
)

func main() {
	fmt.Println("Interactive-MathGo version 1.0")
	inter := mathgo.NewInterpreter()
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	inter.Run(input)
	fmt.Print(input)
}
