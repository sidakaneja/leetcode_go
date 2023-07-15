func sum(nums [] int) int {
    var sum int
    for _, num := range nums {
        sum += num
    }
    return sum
}

func pivotIndex(nums []int) int {
    nums_sum, cur_sum := sum(nums), 0
    
    for i, num := range nums {
        if nums_sum - num == cur_sum {
            return i
        }
        nums_sum -= num
        cur_sum += num
    }
    return -1
}
