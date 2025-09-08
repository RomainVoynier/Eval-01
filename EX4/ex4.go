package main

func Diviseur(a,b int) int {
	if a < 0 {
		return 0
	}
	if b < 0 { 
		return 0
	}
	if a == b {
		return 1
	}
	return 0
	else {
		return a%b
	}
}
