func findMin(nums []int) int {
	var binarySearch func(arr []int, low, high int) int

	binarySearch = func(arr []int, low, high int) int {
		fmt.Println("low is", low, "and high is", high)
		if low == high {
			fmt.Println("return")
			return arr[low]
		}

		mid := low + ((high - low) /2)
		fmt.Println("mid is", mid)
		
		midval := arr[mid]
		fmt.Println("midval is", midval)

		if arr[mid] > arr[high] {
			fmt.Println("go right")
			return binarySearch(arr, mid+1, high)
		} else {
			fmt.Println("go left")
			return binarySearch(arr, low, mid)
		}

		fmt.Println("found min")
		return arr[mid]
	}

	return binarySearch(nums, 0, len(nums)-1)
}
