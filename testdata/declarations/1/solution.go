package declarations

func Declarations() {
	a1 := 1
	a2 := "abc"
	a3, a4, a5 := 1, "abc", int32(a1)

	_, _, _, _, _ = a1, a2, a3, a4, a5
}
