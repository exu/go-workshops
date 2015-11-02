package main

func in(val int) {
	println("in:", val)
}

func out(val int) {
	println("out:", val)
}

func main() {
	in(1)
	defer out(1)
	in(2)
	defer out(2)
	in(3)
	defer out(3)

	in(4)
	in(5)
	in(6)
}
