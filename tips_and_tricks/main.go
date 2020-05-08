package main

import (
	"bufio"
	"fmt"
	"net"
	"regexp"
	"strconv"
)

func main() {
	pattern := regexp.MustCompile(`(?m)What is the answer to (\d+) \+ (\d+)?`)

	conn, err := net.Dial("tcp", "c0e95b4d5ea4ce7e.247ctf.com:50473")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		matches := pattern.FindStringSubmatch(line)
		if len(matches) == 3 {
			x, err := strconv.ParseInt(matches[1], 10, 64)
			if err != nil {
				fmt.Println("Cannot parse int: ", matches[1])
				continue
			}
			y, err := strconv.ParseInt(matches[2], 10, 64)
			if err != nil {
				fmt.Println("Cannot parse int: ", matches[2])
				continue
			}
			fmt.Fprintf(conn, "%d\r\n", x+y)
		}
	}
}
