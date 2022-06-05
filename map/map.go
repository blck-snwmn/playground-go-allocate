package main

//go:noinline
func setMap(m map[int]string) {
	m[1] = "aaaa1"
	m[2] = "aaaa2"
	m[3] = "aaaa3"
	m[4] = "aaaa4"
}

//go:noinline
func returnGenMap() map[int]string {
	return map[int]string{
		1: "aaaa1",
		2: "aaaa2",
		3: "aaaa3",
		4: "aaaa4",
	}
}

var m = map[int]string{
	1: "aaaa1",
	2: "aaaa2",
	3: "aaaa3",
	4: "aaaa4",
}

//go:noinline
func returnGeneratedMap() map[int]string {
	return m
}
