package MovingOn

func InterfaceSliceToStringSlice(target []interface{}) []string {
	var res []string
	for _, value := range target {
		parsedValue, parseSuccess := value.(string)
		if !parseSuccess {
			return nil
		}
		res = append(res, parsedValue)
	}
	return res
}
