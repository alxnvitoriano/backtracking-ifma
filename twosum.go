package main

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		complemento := target - v

		if idx, ok := m[complemento]; ok {
			return []int{idx, i}
		}
		m[v] = i
	}

	return nil // não existe par
}
