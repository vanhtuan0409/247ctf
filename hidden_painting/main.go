package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func hex2int(str string) uint64 {
	cleaned := strings.Replace(str, "0x", "", -1)
	result, _ := strconv.ParseUint(cleaned, 16, 64)
	return result
}

func main() {
	f, err := os.Open("./secret_map.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		x, y := hex2int(parts[0]), hex2int(parts[1])
		fmt.Printf("%d - %d\n", x, y)
	}

}
