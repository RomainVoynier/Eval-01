package main

func PGCD(a,b int) int {
	a := 20
	b := 10
	for !b = 0 {
		a,b = a%b
	}
}
