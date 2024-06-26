package Word_Count

import (
	"os"
	"fmt"
	"bufio"
)

func Count_Words(str string) int {
	fs, err := os.Open(str)

	if err != nil{
		fmt.Println("Error Opening File")
		return 0
	}

	defer fs.Close()

	scanner := bufio.NewScanner(fs)
	scanner.Split(bufio.ScanWords)

	count := 0

	for scanner.Scan() {
		count ++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return 0
	}

	return count
}
