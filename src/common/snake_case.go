package common

import "unicode"



func ConvertToSnakeCase(str string) string{
	var res []rune
	lastChaWasUpper := false

	for _, ch := range str {
		if unicode.IsUpper(ch){
			if !lastChaWasUpper{
				if len(res) > 0 {
					res = append(res, '_')
				}
			}
			lastChaWasUpper = true
		} else {
			lastChaWasUpper = false
		}
		res = append(res, unicode.ToLower(ch))
	}
	return string(res)
}