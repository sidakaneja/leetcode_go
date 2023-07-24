func twoSum(nums []int, target int) []int {
    indices := make(map[int] int)

    for i, num := range nums {
        if val, ok := indices[target - num]; ok {
            return []int{val, i}
        }
        indices[num] = i
    }
    return []int{0,0}
}
