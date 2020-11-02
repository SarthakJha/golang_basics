package main

func main() {
	// long method
	i := 1
	for i <= 10 {
		i++
	}
	println(i)

	// short method
	for j := 1; j < i; j++ {
		println(j)
	}
}
