package main

func EncontrarDuplicado(nums []int) *int {
	set := make(map[int]bool)

	for _, v := range nums {
		if set[v] {
			return &v
		}
		set[v] = true
	}

	return nil // não existe duplicado
}
