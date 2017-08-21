package main

import "fmt"

func main() {
	var n int
	fmt.Println("Hello! Just type a number and i'll show you the spiral martix from 1 to n*n!")
	fmt.Scanf("%d", &n)
	var a [][]int
	var start int
	start = 0
	value := 0
	if n > 1 {
		for k := 0; n > 1 && k < n/2; k++ {
			for i := 0 + k; i < n-k; i++ {
				a = append(a, []int{})
				for j := 0 + k; j < n-k; j++ {
					if k == 0 {
						value = getValue(i, j, n, 0, 0)
						a[i] = append(a[i], value)
					} else {
						value = getValue(i, j, n, start, k)
						a[i][j] = value
					}
				}
			}
			start = start + (n-2*k)*4 - 4
		}

		if n%2 == 1 {
			a[n/2][n/2] = n * n
		}
	} else {
		a = make([][]int, 1, 1)
		a[0] = make([]int, 1, 1)
		a[0][0] = 1
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%3d\n", a[i])
	}

}

func getValue(i int, j int, n int, start int, k int) int {
	var value int
	if i == 0+k || j == n-1-k {
		value = start + i + j + 1 - 2*k
	}

	if i == n-k-1 && i != j {
		value = start + n*2 + n - j - 2 - 5*k
	}

	if j == 0+k && i > 0+k && i < n-1-k {
		value = start + n*3 + n - i - 3 - 7*k
	}
	return value
}
