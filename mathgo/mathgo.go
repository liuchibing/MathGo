package main

import (
	"bufio"
	"fmt"
	"github.com/liuchibing/mathgo"
	"os"
	"flag"
)

const VERSION string = "0.2.0"

//flags
var fInteractive bool

func init() {
        //interactive
        flag.BoolVar(&fInteractive, "interactive", false, "Enable interactive mode")
        flag.BoolVar(&fInteractive, "i", false, "Enable interactive mode (shorthand)")
}

func main() {
	flag.Parse()
	if fInteractive {
		fmt.Println("Interactive-MathGo version", VERSION)
	}
	inter := mathgo.NewInterpreter()
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
