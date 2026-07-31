func findMin(nums []int) int {
	var binarySearch func(arr []int, low, high int) int

	binarySearch = func(arr []int, low, high int) int {
		if low == high {
			return arr[low]
		}

		mid := low + ((high - low) / 2)		
		if arr[mid] > arr[high] {
			return binarySearch(arr, mid+1, high)
		} else {
			return binarySearch(arr, low, mid)
		}

		return arr[mid]
	}

	return binarySearch(nums, 0, len(nums)-1)
}
