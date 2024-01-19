package array

/*func main() {

	arr := []int{1, 2, 3, 4}

	sum := sum_array(arr)

	fmt.Println(sum)

}
*/

func sum_array(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	return sum
}
