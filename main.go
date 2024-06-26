package main

import (
	"os"
	Word_Count "wc_project/cmd"
	"fmt"
)

func main() {
	args := os.Args[1]

	ans := Word_Count.Count_Words(args)

	fmt.Println("The total words are: ", ans)
}
