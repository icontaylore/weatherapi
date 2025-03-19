package main

import "fmt"

func main() {
	n := 39
	finaly := 0
	for n > 9 {
		count := 1
		res := n
		for res > 0 {
			chislo := res % 10
			res /= 10
			count *= chislo
		}
		n = count
		finaly++
	}
	fmt.Println(finaly)
}
