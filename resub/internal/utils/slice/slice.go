package slice

func DeDupe(i []string) []string {
	var o []string

	seen := make(map[string]struct{})
	for _, v := range i {
		if _, b := seen[v]; b {
			continue
		}
		seen[v] = struct{}{}
		o = append(o, v)
	}

	return o
}
