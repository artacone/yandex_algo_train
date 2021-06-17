package main

import "fmt"

func findL(m, k, p, n int) []int {
	minL := k / (m*(p - 1) + n)
	maxL := (k - 1) / (m*(p - 1) + n - 1)
	var possibleL []int
	for l := minL; l <= maxL; l++ {
		if l != 0 && (m*(p - 1) + n - 1)*l + (k - 1) % l == k - 1 {
			possibleL = append(possibleL, l)
		}
	}
	return possibleL
}

// K = (N - 1)*l + (P - 1)*M*L + i
func ambulance(k1, k2, n2, p2, m int) (p1, n1 int) {
	if n2 > m {
		return -1, -1
	}

	var possibleL []int
	if p2 == 1 && n2 == 1 {
		if k1 <= k2 {
			return 1, 1
		} else {
			for l := k2; l <= k1; l++ {
				possibleL = append(possibleL, l)
			}
		}
	} else {
		possibleL = findL(m, k2, p2, n2)
	}
	result := [2]int{-1, -1}
	for _, l := range possibleL {
		i := ((k1 - 1 - (k1 - 1) % l) / l) + 1
		n1 = i % m
		p1 = ((i - n1) / m) + 1
		if n1 == 0 {
			n1 = m
			p1--
		}
		if result == [2]int{-1, -1} {
			result = [2]int{p1, n1}
		} else {
			if p1 != result[0] {
				result[0] = 0
			}
			if n1 != result[1] {
				result[1] = 0
			}
		}
	}
	return result[0], result[1]
}

func main() {
	var k1, k2, n2, p2, m int
	fmt.Scan(&k1, &m, &k2, &p2, &n2)

	fmt.Println(ambulance(k1, k2, n2, p2, m))
}