//find min or smallest element of an array
/* To find the smallest element of an array select the first element of array and then compare this with other elements.
This is a beginner approach, will do it in optimal approach in later sections
*/

package array

/*func main() {
	arr := []int{12, 54, 2, 32, 234}

	min_element := find_min_element(arr)

	fmt.Println(min_element)

}
*/
func find_min_element(arr []int) int {

	if len(arr) == 0 {
		return -1
	}

	min := arr[0]

	for i := 1; i < len(arr); i++ {

		if arr[i] < min {
			min = arr[i]
		}

	}
	return min

}
