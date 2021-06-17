package main

import "fmt"

func main()  {
	var a, b, c, d, e, f float64
	fmt.Scan(&a, &b, &c, &d, &e, &f)

	det := a*d - b*c
	det1 := e*d - f*b
	det2 := a*f - c*e

	if det == 0 {
		if a == 0 && c == 0 {
			if b == 0 && d == 0 {
				if e == 0 && f == 0 {
					fmt.Println(5)
					return
				}
				fmt.Println(0)
				return
			}
			if det1 != 0 {
				fmt.Println(0)
				return
			} else {
				if b != 0 {
					fmt.Println(4, e/b)
					return
				} else {
					fmt.Println(4, f/d)
					return
				}
			}
		}
		if b == 0 && d == 0 {
			if det2 != 0 {
				fmt.Println(0)
				return
			} else {
				if a != 0 {
					fmt.Println(3, e/a)
					return
				} else {
					fmt.Println(3, f/c)
					return
				}
			}
		}
		if a != 0 {
			k := c/a
			if d == k*b && f == k*e {
				fmt.Println(1, -a/b, e/b)
				return
			}
		} else {
			k := a/c
			if b == k*d && e == k*f {
				fmt.Println(1, -c/d, f/d)
				return
			}
		}
		fmt.Println(0)
	} else {
		fmt.Println(2, det1/det, det2/det)
	}
}
