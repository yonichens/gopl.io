// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

//!+
func main() {
	//TODO:练习 1.3： 做实验测量潜在低效的版本和使用了 strings.Join 的版本的运行时间差异。（1.6 节讲解了部分 time 包，11.4节展示了如何写标准测试程序，以得到系统性的性能评测。）
	fmt.Println(strings.Join(os.Args[0:], " "))
}

//!-
