func containsDuplicate(nums []int) bool {
    contains := make(map[int]bool)
    
    for _, num := range nums {
        if contains[num] {
            return true
        }
        contains[num] = true
    }
    return false
}
