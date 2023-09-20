package logger

func mapToZapParams(ex map[ExtraKey]interface{}) []interface{} {
	res := []interface{}{}
	for k, v := range ex {
		res = append(res, k)
		res = append(res, v)
	}
	return res
}
