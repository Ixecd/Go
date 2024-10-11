package main

import "fmt"

func main() {
	s := []int{11,22,33,44}
	//[0, 2) 类似于python3
	s1 := s[0:2]
	for _, value := range s1 {
		fmt.Println("value = ", value)
	}

	t := []int{}
	t = make([]int, 4)

	copy(t, s);
	for _, value := range t {
		fmt.Println("value = ", value)
	}
}