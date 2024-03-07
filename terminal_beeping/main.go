package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()

		if err != nil {
			fmt.Println(err)
			continue
		}

		if char < 48 || char > 57 {
			fmt.Print("The terminal will only beep with an integer from 1 to 9\n\n")
			continue
		}

		num, err := strconv.Atoi(string(char))
		if err != nil {
			fmt.Println(err)
			continue
		}

		for i := 0; i < num; i++ {
			fmt.Print("\a")
			// time.Sleep(time.Second * 2)
		}
	}
}
