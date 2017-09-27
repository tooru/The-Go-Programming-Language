package p05_uniq

func Uniq(ss []string) []string {
	var s string
	var res []string

	for i := 0; i < len(ss); i++ {
		if s != ss[i] {
			res = append(res, ss[i])
			s = ss[i]
		}
	}
	return res
}
