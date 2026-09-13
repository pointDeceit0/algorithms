package leetcode


type Key struct {
	X int
	Y int
}

/*
time  ~ O(n^2)
space ~ O(n^2)
*/
func largestOverlap(img1 [][]int, img2 [][]int) int {

	var d1 []Key
	var d2 []Key

	for i, arr := range img1 {
		for j, v := range arr {
			if v == 1 {
				d1 = append(d1, Key{X: i, Y: j})
			}
			if img2[i][j] == 1 {
				d2 = append(d2, Key{X: i, Y: j})
			}
		}
	}

	var max_val int

	shift_vals := make(map[Key]int)

	for _, k1 := range d1 {
		for _, k2 := range d2 {
			curr_key := Key{X: k1.X - k2.X, Y: k1.Y - k2.Y}
			shift_vals[curr_key]++
			max_val = max(max_val, shift_vals[curr_key])
		}
	}

	return max_val
}

/*
time  ~ O(n^2) + O(n^3) => O(n^3)
space ~ O(n^2)

52/60 test cases passed, space ruins it

func LargestOverlap(img1 [][]int, img2 [][]int) int {

	d1 := make(map[Key]int)
	d2 := make(map[Key]int)

	for i, arr := range img1 {
		for j, v := range arr {
			if v == 1 {
				d1[Key{X: i, Y: j}] = 1
			}
			if img2[i][j] == 1 {
				d2[Key{X: i, Y: j}] = 1
			}
		}
	}

	if len(d1) < len(d2) {
		tmp := d2
		d2 = d1
		d1 = tmp
	}

	var max_val int

	for i := range d2 {
		for j := range d1 {
			x_shift := j.X - i.X
			y_shift := j.Y - i.Y
			var tmp int

			for k := range d2 {
				if _, ok := d1[Key{X: k.X + x_shift, Y: k.Y + y_shift}]; ok {
					tmp += 1
				}
			}
			max_val = max(max_val, tmp)
		}
	}

	return max_val
}
*/

/*
time  ~ O(n^6)
space ~ O(n^2)

57/60 test cases passed, it is very demonstrative how additional space could affect on the result

func LargestOverlap(img1 [][]int, img2 [][]int) int {

	var max_val int

	for x1 := range img1 {
		for y1, v1 := range img1[x1] {
			if v1 == 1 {
				for x2 := range img2 {
					for y2, v2 := range img2[x2] {
						var tmp int
						if v2 == 1 {
							for x3 := range img2 {
								for y3, v3 := range img2[x3] {
									if (v3 == 1 &&
										x3 + (x1 - x2) >= 0 && y3 + (y1 - y2) >= 0 &&
										x3 + (x1 - x2) < len(img1) && y3 + (y1 - y2) < len(img1) &&
										img1[x3 + (x1 - x2)][y3 + (y1 - y2)] == 1) {
										tmp += 1
									}
								}
							}
						}
						max_val = max(max_val, tmp)
					}
				}
			}
		}
	}
	
	return max_val
}
*/
