package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		return
	}

	// Writing arguments in a single string
	str := os.Args[1]
	for _, v := range os.Args[2:] {
		str += " " + v
	}

	//2. Checking weather str contain "\n" or not ---> executing the ascii-art
	prev := 'a'
	severallines := false
	for _, v := range str {
		if v == 'n' && prev == '\\' {
			severallines = true
		}
		prev = v
	}
	//3. Writing text line by line into res
	res := ""
	if severallines {
		args := strings.Split(str, "\\n")
		for _, word := range args {
			for i := 0; i < 8; i++ {
				for _, letter := range word {
					res += GetLine(1 + int(letter-' ')*9 + i)
				}
				fmt.Println(res)
				res = ""
			}
		}

	} else {
		for i := 0; i < 8; i++ {
			for _, letter := range str {
				res += GetLine(1 + int(letter-' ')*9 + i)
			}
			fmt.Println(res)
			res = ""
		}
	}
}

func GetLine(num int) string {
	str := ""
	f, e := os.Open("standard.txt")
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(0)
	}
	defer f.Close()

	f.Seek(0, 0)
	content := bufio.NewReader(f)
	for i := 0; i < num; i++ {
		str, _ = content.ReadString('\n')
	}
	str = strings.TrimSuffix(str, "\n")
	return str
}

// func GetLine(num int) string {

// 	f, e := os.Open("standard.txt")
// 	if e != nil {
// 		fmt.Println(e.Error())
// 		os.Exit(0)
// 	}
// 	defer f.Close()

// 	content := bufio.NewScanner(f)
// 	lineNum := 0
// 	line := ""
// 	for content.Scan() {
// 		if lineNum == num {
// 			line = content.Text()
// 		}
// 		lineNum++
// 	}
// 	return line
// }
