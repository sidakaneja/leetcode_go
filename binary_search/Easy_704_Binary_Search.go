// Iterative
func search(nums []int, target int) int {
    lo, hi := 0, len(nums) - 1

    for lo <= hi {
        mid := lo + ((hi - lo) / 2)
        
        if target == nums[mid] {
            return mid
        } else if target < nums[mid] {
            hi = mid - 1
        } else {
            lo = mid + 1
        }
    }

    return -1
}

// Recursive
func binarySearch(nums[] int, target int, i int, j int) int {
    if i > j {
        return -1
    }
    
    mid := i + ((j-i) / 2)
    if target == nums[mid] {
        return mid
    } else if target < nums[mid] {
        return binarySearch(nums, target, i, mid -1)
    } else {
        return binarySearch(nums, target, mid + 1, j)
    }
}
func search(nums []int, target int) int {
    return binarySearch(nums, target, 0, len(nums) - 1)
}
