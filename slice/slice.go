package slice

//go:noinline
func setSlice(b []byte) {
	b[0] = 0x10
	b[1] = 0x11
	b[2] = 0x12
}

//go:noinline
func returnGenSlice() []byte {
	return []byte{0x10, 0x11, 0x12}
}

var dest = []byte{0x10, 0x11, 0x12}

//go:noinline
func returnSlice() []byte {
	return dest
}
