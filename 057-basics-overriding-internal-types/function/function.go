package main

type MyFunc func() string

func call(callback MyFunc) {
	println(callback())
}

func f1() string {
	return "f1"
}

func f2() string {
	return "f2"
}

func f3() string {
	return "f3"
}

func main() {
	call(f2)
	call(f1)
	call(f3)
	call(f2)
}
