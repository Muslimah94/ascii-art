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

    	// 1. Checking is there a new template or not
    	template := "standard.txt"
    	filename := ""
    	args := os.Args[1:]
    	for _, v := range args {
    		if v == "shadow" {
    			template = "shadow.txt"
    		}
    		if v == "thinkertoy" {
    			template = "thinkertoy.txt"
    		}
    		if len(v) > 9 && v[:9] == "--output=" {
    			filename = v[9:]
    		}
    	}

    	// 2. Writing arguments in a single string
    	str := os.Args[1]
    	for _, v := range os.Args[2:] {
    		if v != "standard" && v != "thinkertoy" && v != "shadow" && (len(v) > 9 && v[:9] != "--output=") {
    			str += " " + v
    		}
    	}

    	// 3. Checking weather str contain "\n" or not
    	prev := 'a'
    	severallines := false
    	for _, v := range str {
    		if v == 'n' && prev == '\\' {
    			severallines = true
    		}
    		prev = v
    	}

    	// 4. Creating a file
    	f, e := os.Create(filename)
    	if e != nil {
    		fmt.Println("Please, write the flag \"--output=\", followed by name of the file to create")
    		os.Exit(0)
    	}
    	defer f.Close()
    	res := ""
    	if severallines {
    		args := strings.Split(str, "\\n")
    		for _, word := range args {
    			for i := 0; i < 8; i++ {
    				for _, letter := range word {
    					res += GetLine(template, 1+int(letter-' ')*9+i)
    				}
    				// 5. Writing data into a file
    				fmt.Fprintln(f, res)
    				res = ""
    			}
    		}
    	} else {
    		for i := 0; i < 8; i++ {
    			for _, letter := range str {
    				res += GetLine(template, 1+int(letter-' ')*9+i)
    			}
    			// 5. Writing data into a file
    			fmt.Fprintln(f, res)
    			res = ""
    		}
    	}
    }

    func GetLine(filename string, num int) string {

    	f, e := os.Open(filename)
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