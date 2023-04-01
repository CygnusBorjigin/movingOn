package MovingOn

func StringInterfaceGetKeys(domain map[string]interface{}) []string {
	var res []string
	for key, _ := range domain {
		res = append(res, key)
	}
	return res
}

func StringInterfaceGetValue(domain map[string]interface{}) []interface{} {
	var res []interface{}
	for _, value := range domain {
		res = append(res, value)
	}
	return res
}

func StringInterfaceGetNth(domain map[string]interface{}, target int) (string, interface{}, *string) {
	if target < 0 {
		errorMessage := "Target index negative"
		return "", nil, &errorMessage
	}

	if target > len(domain) {
		errorMessage := "index out of range"
		return "", nil, &errorMessage
	}

	counter := 0

	for key, value := range domain {
		if counter == target {
			return key, value, nil
		}
	}
	errorMessage := "Unexpected error"
	return "", nil, &errorMessage
}

func InterfaceToMapStringInterface(domain interface{}) (map[string]interface{}, bool) {
	parsedRes, parsedSuccess := domain.(map[string]interface{})
	if !parsedSuccess {
		return nil, false
	}
	return parsedRes, true
}

func MapStringInterfaceAccessField(domain map[string]interface{}, target []string) interface{} {
	if len(target) == 1 {
		if !MapStringInterfaceContainKey(domain, target[0]) {
			return nil
		}
		return domain[target[0]]
	}
	currentTarget := target[0]
	if !MapStringInterfaceContainKey(domain, currentTarget) {
		return nil
	}
	currentRes := domain[currentTarget]
	currentResMapStringInterface, isMapStringInterface := currentRes.(map[string]interface{})
	if !isMapStringInterface && len(target) != 1 {
		return nil
	}
	return MapStringInterfaceAccessField(currentResMapStringInterface, target[1:])
}

func StringInterfaceSubMapByIncluding(domain map[string]interface{}, target []string) map[string]interface{} {
	res := make(map[string]interface{})

	for key, value := range domain {
		if StringSliceContainsString(target, key) {
			res[key] = value
		}
	}

	return res
}

func StringInterfaceSubMapByExcluding(domain map[string]interface{}, target []string) map[string]interface{} {
	res := make(map[string]interface{})
	for key, value := range domain {
		if !StringSliceContainsString(target, key) {
			res[key] = value
		}
	}
	return res
}
