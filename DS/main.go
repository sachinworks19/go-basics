package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{23, 45, 78, 98}

	sort.Ints(arr)

	fmt.Println(arr)
}
