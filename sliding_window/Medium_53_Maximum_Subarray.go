func maxSubArray(nums []int) int {
    curSum, maxSum := 0, nums[0]
    for _, num := range nums {
        if curSum < 0 { curSum = 0 }
        curSum += num
        if curSum > maxSum { maxSum = curSum }
    }
    return maxSum
}


// Divide and Conquer
func max(nums ...int) int {
    max := nums[0]
    for _, num := range nums {
        if num > max {
            max = num
        }
    }
    return max
}
func maxSubArray(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    mid := len(nums) / 2
    left := maxSubArray(nums[:mid])
    right := maxSubArray(nums[mid:])

    curSum, maxLeftSum, maxRightSum := 0, 0, 0
    for i := mid - 1; i >= 0; i-- {
        curSum += nums[i]
        maxLeftSum = max(curSum, maxLeftSum)
    } 
    curSum = 0
    for i := mid + 1; i < len(nums); i++ {
        curSum += nums[i]
        maxRightSum = max(curSum, maxRightSum)
    } 

    return max(left, right, maxLeftSum + maxRightSum + nums[mid])
}
