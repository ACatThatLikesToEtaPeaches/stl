package stl


// IsSubSet 判断sub是否是 supper 的一个子集
func IsSubSet(sub []string, supper []string) bool {
	hash := make(map[string]bool)
	for _, w := range supper {
		hash[w] = true
	}
	for _, w := range sub {
		if _, ok := hash[w]; !ok {
			return false
		}
	}
	return true
}
