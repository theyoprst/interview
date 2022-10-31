package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func Sum(plus, minus []string, base int) (*big.Int, bool) {
	sum := big.NewInt(0)
	for _, str := range plus {
		x, success := big.NewInt(0).SetString(str, base)
		if !success {
			return nil, false
		}
		sum.Add(sum, x)
	}
	for _, str := range minus {
		x, success := big.NewInt(0).SetString(str, base)
		if !success {
			return nil, false
		}
		sum.Sub(sum, x)
	}
	return sum, true
}

func Solve(input string) int {
	var leftPlus, leftMinus, rightPlus, rightMinus []string
	wasMinus := false
	wasEq := false
	start := -1
	for i, ch := range input + " " {
		if ch == '-' {
			wasMinus = true
		} else if ch == '+' {
			wasMinus = false
		} else if ch == '=' {
			wasMinus = false
			wasEq = true
		}
		isChDigit := '0' <= ch && ch <= '9' || 'A' <= ch && ch <= 'Z'
		if isChDigit && start == -1 {
			start = i
		}
		if !isChDigit && start != -1 {
			if !wasEq {
				if wasMinus {
					leftMinus = append(leftMinus, input[start:i])
				} else {
					leftPlus = append(leftPlus, input[start:i])
				}
			} else {
				if wasMinus {
					rightMinus = append(rightMinus, input[start:i])
				} else {
					rightPlus = append(rightPlus, input[start:i])
				}
			}
			start = -1
		}
	}
	for base := 2; base <= 100; base++ {
		sumLeft, success := Sum(leftPlus, leftMinus, base)
		if !success {
			continue
		}
		sumRight, success := Sum(rightPlus, rightMinus, base)
		if !success {
			continue
		}
		if sumLeft.Cmp(sumRight) == 0 {
			return base
		}
	}
	return -1
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println(Solve(input))
}
