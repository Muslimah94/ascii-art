package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	if len(os.Args) == 1 {
		return
	}

	font := "standard.txt"
	color := "white"
	str := ""
	sliceData := ""

	// 1. Gathering info (input and flags)
	for i, v := range os.Args[1:] {

		if len(v) >= 8 && v[0:8] == "--color=" {
			color = v[8:]
		} else if len(v) > 8 && v[0:8] == "--slice=" {
			sliceData = v[8:]
		} else if len(v) > 1 && v[0:2] == "--" {
			continue
		} else if v == "standard" || v == "shadow" || v == "thinkertoy" {
			font = v + ".txt"
		} else {
			str += v
			if i != len(os.Args[1:])-1 {
				str += " "
			}

		}
	}

	// 2. Defining slice of string to be colored (if there's one)
	num1, num2, slice := 0, 0, false
	if sliceData != "" {
		num1, num2, slice = getSliceIndexes(sliceData)
	}

	//3. Writing text line by line with necessary color
	args := strings.Split(str, "\\r\\n")
	for _, word := range args {
		for i := 0; i < 8; i++ {
			for j, letter := range word {
				if slice {
					if j >= num1 && j <= num2-1 {
						fmt.Print(ESCseq(color), GetLine(1+int(letter-' ')*9+i, font))
					} else {
						fmt.Print(ESCseq("white"), GetLine(1+int(letter-' ')*9+i, font))
					}
				} else {
					fmt.Print(ESCseq(color), GetLine(1+int(letter-' ')*9+i, font))
				}
			}
			fmt.Printf("\n")
		}
	}
}

// GetLine is a function to read font from files
func GetLine(num int, filename string) string {

	f, e := os.Open("../fonts/" + filename)
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(0)
	}
	scanner := bufio.NewScanner(f)
	lineNum := 0
	line := ""
	for scanner.Scan() {
		if lineNum == num {
			line = scanner.Text()
		}
		lineNum++
	}
	return line
}

// getSliceIndexes is a function that takes data from --slice= flag and converts it into integers
func getSliceIndexes(a string) (int, int, bool) {
	index := -1
	for i, v := range a {
		if v == ':' {
			if i == 0 || i == len(a)-1 {
				fmt.Println("Please select a slice in a right format: <int64:int64>")
				os.Exit(0)
			}
			index = i
			break
		}
	}
	num1, e1 := strconv.Atoi(a[:index])
	num2, e2 := strconv.Atoi(a[index+1:])
	if e1 != nil {
		log.Fatal(e1)
		os.Exit(1)
	} else if e2 != nil {
		log.Fatal(e2)
		os.Exit(1)
	}

	return num1, num2, true
}

// ESCseq is a function that return an esc seqnuence of necessary color
func ESCseq(color string) string {
	//func ESCseq(a, b, c int, color string) string {

	switch color {
	case "white":
		return "\u001b[38;2;255;255;255m"
	case "black":
		return "\u001b[38;2;0;0;0m"
	case "red":
		return "\u001b[38;2;255;0;0m"
	case "green":
		return "\u001b[38;2;0;255;0m"
	case "blue":
		return "\u001b[38;2;0;0;255m"
	case "yellow":
		return "\u001b[38;2;255;255;0m"
	case "pink":
		return "\u001b[38;2;255;0;255m"
	case "grey":
		return "\u001b[38;2;128;128;128m"
	case "purple":
		return "\u001b[38;2;160;32;255m"
	case "brown":
		return "\u001b[38;2;160;128;96m"
	case "orange":
		return "\u001b[38;2;255;160;16m"
	case "cyan":
		return "\u001b[38;2;0;183;235m"
	}

	//return "\u001b[38;2;" + strconv.Itoa(a) + ";" + strconv.Itoa(b) + ";" + strconv.Itoa(c) + "m"
	return "\u001b[38;2;255;255;255m"
}
