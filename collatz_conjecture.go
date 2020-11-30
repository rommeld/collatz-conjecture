package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	//excluding the possibility of a zero or a negative number
	if n <= 0 {

		return 0, errors.New("")
	}
	i := 0
	for ; n > 1; i++ {
		if n%2 == 0 { //calculation for even number
			n /= 2
		} else { //calculation for an odd number
			n = (3 * n) + 1
		}

	}

	return i, nil
}
