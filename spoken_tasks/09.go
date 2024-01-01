package ways

func createMap() {
	var _ map[int]int         // nil
	var _ = make(map[int]int) // initialized, not nil
	var _ = map[int]int{      // composite literal, not nil
		1: 2,
		3: 4,
	}
	var _ = new(map[int]int) // pointer to nil, not usable
}

func createSlice() {
	var _ []int             // nil
	var _ = make([]int, 10) // initialized, not nil
	var _ = []int{1, 2, 3}  // composite literal
	var _ = new([]int)      // pointer to nil, not usable
}
