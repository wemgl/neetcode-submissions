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

		if arr[mid] == target {
			return mid
		} else if arr[low] <= arr[mid] {
			if arr[low] <= target && target < arr[mid] {
				return binarySearch(arr, low, mid-1, target)
			} else {
				return binarySearch(arr, mid+1, high, target)
			}
		} else {
			if arr[mid] < target && target <= arr[high] {
				return binarySearch(arr, mid+1, high, target)
			} else {
				return binarySearch(arr, low, mid-1, target)
			}
		}

		return mid
	}

	return binarySearch(nums, 0, len(nums)-1, target)
}