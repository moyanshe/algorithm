package classics

import "strings"

//ReverseWords 反转单词
func ReverseWords(s string) string {
	strList := strings.Split(s, " ")
	var res []string
	for i := len(strList) - 1; i >= 0; i-- {
		str := strings.TrimSpace(strList[i])
		if len(str) > 0 {
			res = append(res, strList[i])
		}
	}
	return strings.Join(res, " ")
}
