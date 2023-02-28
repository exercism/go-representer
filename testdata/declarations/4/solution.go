package declarations

func Declarations() {
	var (
		a1         int
		a2         int32
		a3, a4     int16
		a5                 = 3
		a6                 = uint16(4)
		a7, a8, a9         = 5, int16(2), a2
		a10, a11   float32 = 1, 3.4
	)
	_, _, _, _, _, _, _, _, _, _, _ = a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11
}
