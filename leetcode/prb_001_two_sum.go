package leetcode

func twoSum(nums []int, target int) []int {

    d := make(map[int]int)

	for i := range nums {
		if val, ok := d[target - nums[i]]; ok {
			return []int{val, i}
		}
		d[nums[i]] = i
	}
	return []int{}
}