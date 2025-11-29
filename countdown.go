package main

func ContarParesDistintos(nums []int) int {
	set := make(map[int]bool)

	for _, v := range nums {
		set[v] = true
	}

	n := len(set)
	return (n * (n - 1)) / 2
}
