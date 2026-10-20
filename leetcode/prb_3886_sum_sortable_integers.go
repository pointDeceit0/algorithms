package leetcode


func sortableIntegers(nums []int) int {
    n := len(nums)
	res := 0

	for k := 1; k <= n; k++ {
		if n % k != 0 {
			continue
		}

		isPossible := true
		var prevMax int

		for i := 0; i < n; i += k {
			block := nums[i : i+k]
			
			minVal, maxVal := block[0], block[0]
			drops := 0

			for j := 0; j < k; j++ {
				if block[j] < minVal {
					minVal = block[j]
				}
				if block[j] > maxVal {
					maxVal = block[j]
				}

				if j < k-1 && block[j] > block[j+1] {
					drops++
				}
			}

			if drops > 1 || (drops == 1 && block[k-1] > block[0]) {
				isPossible = false
				break
			}

			if i > 0 && minVal < prevMax {
				isPossible = false
				break
			}

			prevMax = maxVal
		}

		if isPossible {
			res += k
		}
	}

	return res
}