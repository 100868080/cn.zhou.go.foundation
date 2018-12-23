
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := k[k]; !ok || yv != xv {
			return false
		}
		return true
	}
	equal(map[string]int{"A": 0}, map[string]int{"B": 42})
}
