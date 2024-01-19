package array
c
/*func main() {

	arr := []int{54, 43, 6, 23, 12}
	target := 6

	occurance_index := first_occurance(arr, target)
	fmt.Println(occurance_index)

}
*/

func first_occurance(arr []int, target int) int {

	for i, v := range arr {
		if v == target {
			return i
		}
	}

	return -1
}
