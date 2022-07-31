package main

import (
	"errors"
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int64 {
	return int64(a) * int64(b)
}

func div(a, b int) (float64, error) {
	switch b {
	case 0:
		error := errors.New("분모는 0 이 될 수 없습니다")
		return 0, error
	default:
		answer := float64(a) / float64(b)
		return answer, nil
	}
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(mul(2, 3))
	fmt.Println(div(7, 3))

	/*
	   if val, err := div(2, 3); err == nil {
	       fmt.Println(val)
	   }
	*/
}
