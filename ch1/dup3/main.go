// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 12.

//!+

// Dup3 prints the count and text of lines that
// appear more than once in the named input files.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	//简化：由于 ReadFile 函数需要文件名作为参数，因此只读指定文件，不读标准输入
	for _, filename := range os.Args[0:] {
		data, err := ioutil.ReadFile(filename)
		fmt.Printf("%s\n", os.Args[0:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		//简化：由于行计数代码只在一处用到，故将其移回	 main 函数
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

//!-
