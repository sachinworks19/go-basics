//find max or largest element of an array
/* To find the largest element of an array select the first element of array and then compare this with other elements.
This is a beginner approach, will do it in optimal approach in later sections
*/

package array

/*func main() {

	arr := []int{21, 12, 123, 43, 34, 66}

	max_element := find_max_element(arr)

	fmt.Println("Largest Element in the array is ", max_element)

}
*/

func find_max_element(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	max := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}

	}

	return max
}
