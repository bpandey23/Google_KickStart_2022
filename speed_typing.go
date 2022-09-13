package main

import (
	//"strings"
	"fmt"
)

func main() {
	var T int
	_, err := fmt.Scan(&T)
	if err != nil {
		fmt.Println("error in scan")
	}

	for i := 0; i < T; i++ {
		var in, p string
		fmt.Scan(&in)
		fmt.Scan(&p)
		re, ans := compare(in, p)
		if re {
			fmt.Println(fmt.Sprintf("Case #%v: %v", i+1, ans))

		} else {
			fmt.Println(fmt.Sprintf("Case #%v: IMPOSSIBLE", i+1))

		}

	}

}

func compare(in, p string) (bool, int) {
	in_b := []byte(in)
	p_b := []byte(p)
	l_in := len(in_b)
	l_p := len(p_b)
	i, j := 0, 0
	for i, j = 0, 0; j < l_in && i < l_p; i++ {
		if in_b[j] == p_b[i] {
			j++
		}
	}
	if j == l_in {
		fmt.Println("j value ", j, l_in)
		return true, l_p - j
	}
	fmt.Println("j value ", j, l_in)
	return false, -1
}
