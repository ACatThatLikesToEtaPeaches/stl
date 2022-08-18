package stl

func ReverseString(origin string) string {
	res := make([]rune, len(origin))
	for i, c := range origin {
		res[len(origin)-i-1] = c
	}
	return string(res)
}