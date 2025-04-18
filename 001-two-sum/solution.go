func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)

    for i := 0; i < len(nums); i++ {
        complement := target - nums[i]
        if idx, ok := seen[complement]; ok {
            return []int{idx, i}
        }
        seen[nums[i]] = i
    }

    return []int{}
}
