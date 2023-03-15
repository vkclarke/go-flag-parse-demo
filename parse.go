package main

import "strings"

func parseFlags(s []string) ([]string, []string) {
	var flags, vals []string
	for _, v := range s {
		if v[0] != '-' {
			vals = append(vals, v)
			continue
		}
		switch l := len(v); l {
		case 1:
			vals = append(vals, v)
		case 2, 3:
			if v[1] == '-' {
				vals = append(vals, v)
				break
			}
			for j := 1; j < l; j++ {
				flags = append(flags, v[j:j+1])
			}
		default:
			if v[0] == '-' {
				if v[1] == '-' {
					f := v[2:]
					if k := strings.IndexRune(f, '='); k != -1 {
						if len(f[k:]) > 1 {
							vals = append(vals, f[k+1:])
						}
						f = f[:k]
					}
					flags = append(flags, f)
					break
				}
				for j := 1; j < l; j++ {
					flags = append(flags, v[j:j+1])
				}
				break
			}
			vals = append(vals, v)
		}
	}
	return flags, vals
}
