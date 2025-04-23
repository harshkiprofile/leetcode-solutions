func containsNearbyDuplicate(nums []int, k int) bool {
    mFound := make(map[int]int)

    for i, num := range nums {
        if j,ok := mFound[num]; ok && cAbs(i-j) <= k{
            return true
        }
        mFound[num] = i
    }
    return false
}

func cAbs(n int) int {
    if n<0 {
        return -n
    } else {
        return n
    }
}