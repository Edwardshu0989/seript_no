package util

import "sort"

// 按照字母顺序对map的键排序，然后对值进行相加
func Sign(params map[string]string) string {
	var str string
	ks := make([]string, 0, len(params))
	for k := range params {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	for _, k := range ks {
		v := params[k]
		if v == "" {
			continue
		}
		str += v
	}
	return str
}
