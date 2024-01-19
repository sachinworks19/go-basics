package array

/*func main() {

	arr := []int{24, 31, 13, 12}

	second_largest := second_largest_element(arr)

	fmt.Println(second_largest)

}
*/

func second_largest_element(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	large1 := 0
	large2 := 0

	large1 = arr[0]

	for i := 1; i < len(arr); i++ {
		if large1 < arr[i] {
			large2 = large1
			large1 = arr[i]
		} else if large2 < arr[i] {
			large2 = arr[i]
		}
	}
	return large2

}
