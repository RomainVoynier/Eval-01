package main

func main(s string) bool {
	if len(s) < 1 {
		return false
	}
	return main(s[0]) || main(s[1:])
}

func main(a byte) bool {
	return (a <= 'z' && a >= 'a') || (a <= 'Z' && a >= 'A')
}
