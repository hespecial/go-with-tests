package main

func Sum(nums []int) (sum int) {
	for _, v := range nums {
		sum += v
	}
	return
}

func SumAll(nums ...[]int) (sum []int) {
	for _, v := range nums {
		sum = append(sum, Sum(v))
	}
	return
}
