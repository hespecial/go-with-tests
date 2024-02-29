package main

func Sum(nums []int) (sum int) {
	for _, v := range nums {
		sum += v
	}
	return
}

func SumAllTails(nums ...[]int) (sums []int) {
	for _, v := range nums {
		if len(v) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(v[1:]))
		}
	}
	return
}
