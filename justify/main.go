package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var termWidth int

func init() {
	termWidth = GetTermWidth()
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	font := "standard.txt"
	align := "left"
	text := ""

	// 1. Gathering info (input and flags)
	for _, v := range os.Args[1:] {

		if len(v) >= 8 && v[0:8] == "--align=" {
			align = v[8:]
		} else if v == "standard" || v == "shadow" || v == "thinkertoy" {
			font = v + ".txt"
		} else {
			text += v
		}
	}
	if align == "left" || align == "right" || align == "center" || align == "centre" {
		RightLeft(text, font, align)
	} else if align == "justify" {

		lines := strings.Split(text, "\\n")
		for _, v := range lines {
			Justify(v, font)
		}
	}
}

func RightLeft(text, font, align string) {
	res := ""
	args := strings.Split(text, "\\n")
	for _, word := range args {
		for i := 0; i < 8; i++ {
			for _, letter := range word {
				res += GetLine(1+int(letter-' ')*9+i, font)
			}
			if align == "left" {
				fmt.Println(res)
			}
			if align == "right" {
				fmt.Print(printSpaces(termWidth - len(res)))
				fmt.Print(res)
			}
			if align == "center" || align == "centre" {
				fmt.Print(printSpaces((termWidth - len(res)) / 2))
				fmt.Print(res)
				fmt.Print(printSpaces((termWidth - len(res)) / 2))
				fmt.Println()
			}
			res = ""
		}
	}
}

func Justify(words string, font string) {
	sws := SplitWhiteSpacesAWESOME(words)
	ar := make([][]string, len(sws))
	j := 0
	container := ""
	for i := 0; i < 8; i++ {
		j = 0
		for _, letter := range words {
			if letter != ' ' {
				container += GetLine(1+int(letter-' ')*9+i, font)
			} else if letter == ' ' && container != "" {
				ar[j] = append(ar[j], container)
				container = ""
				j++
			}
		}
		ar[j] = append(ar[j], container)
		container = ""
	}
	textLen := 0
	for _, arOfStr := range ar {
		textLen += len(arOfStr[0])
	}
	if len(sws) == 1 {
		RightLeft(words, font, "left")
		return
	}
	numSpaces := (termWidth - textLen) / (len(sws) - 1)

	for i := 0; i < 8; i++ {
		for k, v := range ar {
			fmt.Print(v[i])
			if k != len(ar)-1 {
				fmt.Print(printSpaces(numSpaces))
			}
		}
		fmt.Println()
	}

}

func printSpaces(num int) string {
	a := ""
	for i := 1; i <= num; i++ {
		a += " "
	}
	return a
}

func GetTermWidth() int {
	out, er1 := exec.Command("tput", "cols").Output()
	out1 := strings.TrimSuffix(string(out), "\n")
	num, er2 := strconv.Atoi(string(out1))
	if er1 != nil {
		log.Fatal(er1)
		os.Exit(1)
	}
	if er2 != nil {
		log.Fatal(er2)
		os.Exit(1)
	}
	return num
}

func GetLine(num int, filename string) string {

	f, e := os.Open("../fonts/" + filename)
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(0)
	}
	defer f.Close()

	content := bufio.NewScanner(f)
	lineNum := 0
	line := ""
	for content.Scan() {
		if lineNum == num {
			line = content.Text()
		}
		lineNum++
	}
	return line

}

func SplitWhiteSpacesAWESOME(str string) []string {

	nbw := 0
	var prel rune
	prel = ' '
	word := ""

	for _, v := range str {
		if (v != ' ' && v != '\t' && v != '\n') && (prel == ' ' || prel == '\t' || prel == '\n') {
			nbw++
		}
		prel = v
	}

	ar := make([]string, nbw)
	a := 0
	for _, v := range str {
		if v != ' ' && v != '\t' && v != '\n' {
			word = word + string(v)
		}
		if v == ' ' || v == '\t' || v == '\n' && word != "" {
			a++
			ar[a-1] = word
			word = ""
		}

	}

	if word != "" {
		ar[a] = word
	}
	return ar
}
