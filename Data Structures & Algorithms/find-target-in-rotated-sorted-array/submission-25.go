func search(nums []int, target int) int {
	var binarySearch func(arr []int, low, high, target int) int

	binarySearch = func(arr []int, low, high, target int) int {
		if low == high {
			if arr[low] == target {
				return low
			}
			return -1
		}

		mid := low + ((high - low)/2)
		fmt.Println("low =", low, "high =", high, "mid =", mid, "arr[low] =", arr[low], "arr[mid] =", arr[mid], "arr[high] =", arr[high])

		if arr[mid] == target {
			return mid
		} else if arr[low] <= arr[mid] {
			if arr[low] <= target && target < arr[mid] {
				fmt.Println("go left 1")
				return binarySearch(arr, low, mid-1, target)
			} else {
				fmt.Println("go right 1")
				return binarySearch(arr, mid+1, high, target)
			}
		} else {
			if arr[mid] < target && target <= arr[high] {
				fmt.Println("go right 2")
				return binarySearch(arr, mid+1, high, target)
			} else {
				fmt.Println("go left 2")
				return binarySearch(arr, low, mid-1, target)
			}
		}

		return mid
	}

	return binarySearch(nums, 0, len(nums)-1, target)
}