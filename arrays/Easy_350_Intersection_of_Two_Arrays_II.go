// using sort
func intersect(nums1 []int, nums2 []int) []int {
    // Sorting
    sort.Ints(nums1)
    sort.Ints(nums2)
    
    result := []int{}
    i, j := 0, 0

    for i < len(nums1) && j < len(nums2) {
        if nums1[i] < nums2[j] {
            i++
        } else if nums2[j] < nums1[i] {
            j++
        } else {
            result = append(result, nums1[i])
            i++
            j++
        }
    }
    return result
}

// using a map
func intersect(nums1 []int, nums2 []int) []int {
    nums1Counts := make(map[int]int)
    
    for _, num := range nums1 {
        nums1Counts[num]++
    }
    result := []int{}

    for _, num := range nums2 {
        if nums1Counts[num] > 0 {
            result = append(result, num)
            nums1Counts[num]--
        }
    }
    return result
}
