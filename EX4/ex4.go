package main

func PGCD(a,b int) int {
	a = 20
	b = 10
	if a < 0 {
		return 0
	}
	if b < 0 { 
		return 0
	}
	if a == b {
		return 1
	}
	else {
		return a%b
	}
}
