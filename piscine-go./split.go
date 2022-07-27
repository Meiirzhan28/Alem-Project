package piscine

func Split(s, sep string) []string {
	s = s + sep
	cut := 0
	var res []string
	for i := 0; i < len(s); i++ {
		if Index(s, sep) == i {
			temp := s[cut:i]
			if temp != "" {
				res = append(res, temp)
			}
			s = s[i+len(sep):]
			i = 0
		}
	}
	return res
}
