package main

func Sum(nums []int) (sum int) {
	for _, v := range nums {
		sum += v
	}
	return
}

func SumAllTails(nums ...[]int) (sum []int) {
	for _, v := range nums {
		sum = append(sum, Sum(v[1:]))
	}
	return
}
