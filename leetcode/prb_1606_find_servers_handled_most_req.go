package leetcode


func busiestServers(k int, arrival []int, load []int) []int {
	// Brute Force ~ O(nm) -- TLE 105/108

    m := 0
	servers := make([]int, k)
	done := make([]int, k)
	deltas := make([]int, k)
	for i, a := range arrival {

		cnt, j := 0, i % k
		for cnt < k {		
			if max(servers[j] - (a - deltas[j]), 0) == 0 {
				servers[j] = load[i]
				deltas[j] = a
				done[j]++
				if done[j] > m {
					m = done[j]
				}
				break
			} else {
				servers[j] = servers[j] - (a - deltas[j])
				deltas[j] = a
			}
			cnt++; j = (j + 1) % k
		}
	}

	ans := []int{}
	for i, d := range done {
		if d == m {
			ans = append(ans, i)
		}
	}

	return ans
}