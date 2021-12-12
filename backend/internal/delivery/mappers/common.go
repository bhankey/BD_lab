package mappers

func SliceInt64ToInt(s []int64) []int {
	r := make([]int, 0, len(s))

	for _, i := range s {
		r = append(r, int(i))
	}

	return r
}
