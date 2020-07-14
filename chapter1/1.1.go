package chapter1

// エラトステネスの篩
func SieveOfEratosthenes(x int) []int {
	remove := func(l []int, i int) []int {
		if i >= len(l) {
			return l
		}

		return append(l[:i], l[i+1:]...)
	}

	a := make([]int, 0, x)
	for i := 2; i <= x; i++ {
		a = append(a, i)
	}

	for _, p := range a {
		if p*p > x {
			break
		}

		for i, v := range a {
			if v <= p {
				continue
			}

			if v%p == 0 {
				a = remove(a, i)
			}
		}
	}

	return a
}
