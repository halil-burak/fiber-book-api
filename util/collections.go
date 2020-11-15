package util

// Filter filters a given slice of string with a predicate boolean returning function
func Filter(strs []string, f func(string) bool) []string {
	fstrs := make([]string, 0)
	for _, v := range strs {
		if f(v) {
			fstrs = append(fstrs, v)
		}
	}
	return fstrs
}
